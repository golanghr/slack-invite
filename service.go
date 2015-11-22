/*
Copyright (c) 2015 Golang Croatia
All rights reserved.

The MIT License (MIT)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Package main ...
package main

import (
	"errors"
	"syscall"

	"github.com/golanghr/platform/handlers"
	"github.com/golanghr/platform/logging"
	"github.com/golanghr/platform/managers"
	"github.com/golanghr/platform/options"
	"github.com/golanghr/platform/servers"
	"github.com/golanghr/platform/services"
	pb "github.com/golanghr/slack-invite/protos"
	"google.golang.org/grpc"
)

// Service - tbd ...
type Service struct {

	// Options - Is a global service options
	options.Options

	// Servicer - Is a Servicer interface
	services.Servicer

	// Grpc - Here GRPC server is located.
	Grpc servers.Serverer

	// HTTP - Here HTTP server is located.
	HTTP servers.Serverer

	// Managerer - Service runtime manager. Manager actually contains start, stop
	// and all of the runtime related management handlers.
	managers.Managerer

	// Logging -
	*logging.Entry

	// Slack -
	*Slack
}

// GrpcServer - Will return back actual grpc.Server
// I understand that this looks like a hack but I'd rather have it require to satisfy interface
// than having it require to satisfy nil.
func (s *Service) GrpcServer() *grpc.Server {
	return s.Grpc.Interface().(*servers.Grpc).Server
}

// Terminate - Will send SIGTERM towards service interrupt signal resulting entire
// service to go down
func (s *Service) Terminate() {
	s.GetInterruptChan() <- syscall.SIGTERM
}

// NewService -
func NewService(opts options.Options, logger *logging.Entry) (*Service, error) {

	serv, err := services.New(opts)

	if err != nil {
		return nil, err
	}

	slackToken, ok := opts.Get("slack-token")

	if !ok {
		return nil, errors.New("You need to provide slack token in order to start service...")
	}

	slackApiDebug, ok := opts.Get("slack-api-debug")

	if !ok {
		return nil, errors.New("You need to provide `slack-api-debug` in order to start service...")
	}

	grpcServer, err := servers.NewGrpcServer(serv, opts, logger)

	if err != nil {
		return nil, err
	}

	httpServer, err := servers.NewHTTPServer(serv, opts, logger)

	if err != nil {
		return nil, err
	}

	serviceManager, err := managers.New(serv, opts, logger)

	if err != nil {
		return nil, err
	}

	// We are about to attach GRPC service now ...
	if err := serviceManager.Attach("grpc", grpcServer); err != nil {
		return nil, err
	}

	// We are about to attach HTTP service now ...
	if err := serviceManager.Attach("http", httpServer); err != nil {
		return nil, err
	}

	sc := &Service{
		Options:   opts,
		Servicer:  serv,
		Grpc:      grpcServer,
		HTTP:      httpServer,
		Entry:     logger,
		Managerer: serviceManager,
		Slack:     NewSlack(slackToken.String(), slackApiDebug.Bool()),
	}

	pb.RegisterSlackServer(sc.GrpcServer(), sc)

	hander, err := handlers.NewHttpGrpcHandler(serv, logger, pb.RegisterSlackHandler)

	if err != nil {
		return nil, err
	}

	httpServer.Interface().(*servers.HTTP).Handler = hander
	return sc, nil
}
