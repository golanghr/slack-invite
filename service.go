// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package main ...
package main

import (
	"time"

	"github.com/golanghr/platform/config"
	"github.com/golanghr/platform/server"
	"github.com/golanghr/platform/service"
)

// SlackInvite -
type SlackInvite struct {
	service.Service

	Config config.Manager

	Server    *server.Server
	ServerTLS *server.Server

	Quit chan bool
}

// LoadConfiguration - Will basically ensure that service defaults are set and that
// they are properly syncronized against etcd
func (si *SlackInvite) LoadConfiguration(cnf map[string]interface{}) error {
	log.Debug("Ensuring configuration ...")

	cm := map[string]string{
		"service-name":        ServiceName,
		"service-description": ServiceDescription,
		"service-version":     ServiceVersion,
		"use-tls":             ServiceUseTLS,
		"addr":                ServiceAddr,
		"timeout":             ServiceServerTimeout,
	}

	if err := si.Config.EnsureSetMany(cm); err != nil {
		return err
	}

	log.Debug("Syncronizing configuration ...")

	// if err := si.Manager.SyncMap(cm); err != nil {
	// 	return err
	// }

	return nil
}

// LoadWebServer -
func (si *SlackInvite) LoadWebServer() (err error) {
	log.Debug("Preparing HTTP/TLS server ...")

	var addr string
	var timeout time.Duration

	if addr, err = si.Config.GetString("addr"); err != nil {
		return
	}

	if timeout, err = si.Config.GetDuration("timeout"); err != nil {
		return
	}

	handler := server.Handler{}

	si.Server = server.NewServer(si, &server.Options{
		Timeout: timeout * time.Second,
		Addr:    addr,
	}, handler)

	si.ServerTLS = server.NewServerTLS(si, &server.Options{
		Timeout: 10 * time.Second,
		Addr:    addr,
	}, handler)

	return
}

// Run -
func (si *SlackInvite) Run() (err error) {
	log.Warn("Starting service ...")

	// Start HTTP and TLS separatedly instead of Server.Start()
	// Use locking here so we wait for http to start, wait for tls to start
	// and than we consider it as done so error or really anything else

	if err = si.Server.ListenAndServe(); err != nil {
		return
	}

	if err = si.ServerTLS.ListenAndServe(); err != nil {
		return
	}

	return
}

// Recover - Will just capture recover error and log it as fatal.
// we will still kill the service as it paniced and service should never ever
// panic and be considered as "ok"
func (si *SlackInvite) Recover() {
	if err := recover(); err != nil {
		LogFatalError(err.(error), "Panic happen! Killing service now...")
	}
}

// NewSlackInvite -
func NewSlackInvite(cnf map[string]interface{}) (*SlackInvite, error) {
	var serv service.Service
	var err error
	var conf config.Manager

	if conf, err = config.New(cnf); err != nil {
		return nil, err
	}

	if serv, err = service.New(conf, logger); err != nil {
		return nil, err
	}

	slackinvite := &SlackInvite{
		Service: serv,
		Config:  conf,
		Quit:    make(chan bool),
	}

	if err := slackinvite.LoadConfiguration(cnf); err != nil {
		return nil, err
	}

	if err := slackinvite.LoadWebServer(); err != nil {
		return nil, err
	}

	return slackinvite, nil
}
