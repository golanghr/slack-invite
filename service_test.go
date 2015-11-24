// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package main ...
package main

import (
	"fmt"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/golanghr/platform/options"
	"github.com/golanghr/platform/utils"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/golanghr/slack-invite/protos"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	alreadyInTeamEmail = "nevio.vesic@gmail.com"
)

func getSlackOptions() (options.Options, error) {
	return options.New("memo", map[string]interface{}{
		"service-name":        utils.GetFromEnvOr("SLACK_SERVICE_NAME", "Slack Invite"),
		"service-description": utils.GetFromEnvOr("SLACK_SERVICE_DESCRIPTION", "Golang.hr Slack Invite is a small automation service written on top of Golang.hr Platform."),
		"service-version":     getFloat(utils.GetFromEnvOr("SLACK_SERVICE_VERSION", "0.1")),
		"formatter":           "text",
		"level":               logrus.DebugLevel,
		"manager-interrupt-wait-timeout": getInt(utils.GetFromEnvOr("SLACK_SERVICE_MANAGER_INTERRUPT_TIMEOUT", "10")),
		"grpc-listen-forever":            getBool(utils.GetFromEnvOr("SLACK_SERVICE_GRPC_LISTEN_FOREVER", "true")),
		"grpc-addr":                      utils.GetFromEnvOr("SLACK_SERVICE_GRPC_ADDR", ":4772"),
		"grpc-tls":                       getBool(utils.GetFromEnvOr("SLACK_SERVICE_GRPC_TLS", "true")),
		"grpc-tls-cert":                  utils.GetFromEnvOr("HELLO_SERVICE_GRPC_TLS_CERT", "test_data/server.crt"),
		"grpc-tls-key":                   utils.GetFromEnvOr("HELLO_SERVICE_GRPC_TLS_KEY", "test_data/server.key"),
		"grpc-tls-domain":                utils.GetFromEnvOr("SLACK_SERVICE_GRPC_TLS_DOMAIN", "golang.hr"),
		"http-addr":                      utils.GetFromEnvOr("SLACK_SERVICE_HTTP_ADDR", ":8500"),
		"http-listen-forever":            getBool(utils.GetFromEnvOr("SLACK_SERVICE_HTTP_LISTEN_FOREVER", "true")),
		"slack-team-name":                utils.GetFromEnvOr("SLACK_TEAM_NAME", "golanghr"),
		"slack-token":                    utils.GetFromEnvOr("SLACK_TOKEN", ""),
		"slack-api-debug":                getBool(utils.GetFromEnvOr("SLACK_API_DEBUG", "false")),
	})
}

func getSlackService(opts options.Options) (*Service, error) {
	return NewService(opts, logger)
}

func TestSlackInviteService(t *testing.T) {
	opts, _ := getSlackOptions()
	service, err := getSlackService(opts)

	Convey("Should be service without any errors", t, func() {
		So(service, ShouldHaveSameTypeAs, &Service{})
		So(err, ShouldBeNil)
	})
}

func TestStatsPb(t *testing.T) {
	opts, _ := getSlackOptions()

	// We gotta update these as there are multiple tests spawning out multiple listeners
	opts.Set("http-addr", ":8501")
	opts.Set("grpc-addr", ":4773")

	service, err := getSlackService(opts)

	Convey("Should be service without any errors", t, func() {
		So(service, ShouldHaveSameTypeAs, &Service{})
		So(err, ShouldBeNil)
	})

	Convey("Test Fucking protobuff", t, func() {
		pb, err := service.Slack.GetStatsPb()
		So(err, ShouldBeNil)

		fmt.Printf("Pb response: %q \n", pb)
	})
}

func TestInviteEndpoint(t *testing.T) {
	opts, _ := getSlackOptions()

	// We gotta update these as there are multiple tests spawning out multiple listeners
	opts.Set("http-addr", ":8502")
	opts.Set("grpc-addr", ":4774")

	service, err := getSlackService(opts)

	Convey("Should be service without any errors", t, func() {
		So(service, ShouldHaveSameTypeAs, &Service{})
		So(err, ShouldBeNil)
	})

	go func() { service.Start() }()
	defer service.Terminate()
	address, _ := service.Options.Get("grpc-addr")

	var gopts []grpc.DialOption

	var creds credentials.TransportAuthenticator
	domain, _ := opts.Get("grpc-tls-domain")
	creds, _ = credentials.NewClientTLSFromFile("test_data/server.crt", domain.String())
	gopts = append(gopts, grpc.WithTransportCredentials(creds))

	Convey("By passing invalid first name we get error that one must be set", t, func() {
		conn, err := grpc.Dial(address.String(), gopts...)
		So(err, ShouldBeNil)
		defer conn.Close()

		client := pb.NewSlackClient(conn)

		_, err = client.Invite(context.Background(), &pb.Request{})
		So(err.Error(), ShouldContainSubstring, "First name must be provided in order")
	})

	Convey("By passing invalid last name we get error that one must be provided", t, func() {
		conn, err := grpc.Dial(address.String(), gopts...)
		So(err, ShouldBeNil)
		defer conn.Close()

		client := pb.NewSlackClient(conn)

		_, err = client.Invite(context.Background(), &pb.Request{FirstName: "Test"})
		So(err.Error(), ShouldContainSubstring, "Last name must be provided in order")
	})

	Convey("By passing invalid email we get error that one must be set", t, func() {
		conn, err := grpc.Dial(address.String(), gopts...)
		So(err, ShouldBeNil)
		defer conn.Close()

		client := pb.NewSlackClient(conn)

		_, err = client.Invite(context.Background(), &pb.Request{FirstName: "Test", LastName: "Testing"})
		So(err.Error(), ShouldContainSubstring, "Valid email must be provided")

		_, err = client.Invite(context.Background(), &pb.Request{FirstName: "Test", LastName: "Testing", Email: "iamnotvalid"})
		So(err.Error(), ShouldContainSubstring, "Valid email must be provided")
	})

	Convey("By passing already invited email we get already invited error", t, func() {
		conn, err := grpc.Dial(address.String(), gopts...)
		So(err, ShouldBeNil)
		defer conn.Close()

		client := pb.NewSlackClient(conn)

		_, err = client.Invite(context.Background(), &pb.Request{FirstName: "Test", LastName: "Testing"})
		So(err.Error(), ShouldContainSubstring, "Valid email must be provided")

		_, err = client.Invite(context.Background(), &pb.Request{FirstName: "Test", LastName: "Testing", Email: alreadyInTeamEmail})
		So(err.Error(), ShouldContainSubstring, "Failed to invite to team: already_in_team")
	})
}
