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
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/golanghr/platform/utils"
)

var (
	serviceOptions = map[string]interface{}{
		"service-name":        utils.GetFromEnvOr("HELLO_SERVICE_NAME", "Slack Invite"),
		"service-description": utils.GetFromEnvOr("HELLO_SERVICE_DESCRIPTION", "Golang.hr Slack Invite is a small automation service written on top of Golang.hr Platform."),
		"service-version":     getFloat(utils.GetFromEnvOr("HELLO_SERVICE_VERSION", "0.1")),
		"formatter":           "text",
		"level":               logrus.DebugLevel,
		"manager-interrupt-wait-timeout": getInt(utils.GetFromEnvOr("SLACK_SERVICE_MANAGER_INTERRUPT_TIMEOUT", "10")),
		"grpc-listen-forever":            getBool(utils.GetFromEnvOr("SLACK_SERVICE_GRPC_LISTEN_FOREVER", "true")),
		"grpc-addr":                      utils.GetFromEnvOr("SLACK_SERVICE_GRPC_ADDR", ":4772"),
		"grpc-tls":                       getBool(utils.GetFromEnvOr("SLACK_SERVICE_GRPC_TLS", "true")),
		"grpc-tls-domain":                utils.GetFromEnvOr("SLACK_SERVICE_GRPC_TLS_DOMAIN", "golang.hr"),
		"grpc-rest-addr":                 utils.GetFromEnvOr("SLACK_SERVICE_REST_ADDR", ":8500"),
		"grpc-rest-listen-forever":       getBool(utils.GetFromEnvOr("SLACK_SERVICE_REST_LISTEN_FOREVER", "true")),
		"http-addr":                      utils.GetFromEnvOr("SLACK_SERVICE_HTTP_ADDR", ":8800"),
		"http-listen-forever":            getBool(utils.GetFromEnvOr("SLACK_SERVICE_HTTP_LISTEN_FOREVER", "true")),
		"slack-team-name":                utils.GetFromEnvOr("SLACK_TEAM_NAME", "golanghr"),
		"slack-token":                    utils.GetFromEnvOr("SLACK_TOKEN", ""),
		"slack-api-debug":                getBool(utils.GetFromEnvOr("SLACK_API_DEBUG", "false")),
		"google-analytics-tracking-id":   utils.GetFromEnvOr("SLACK_GA_TRACKING_ID", ""),
	}
)

func getBool(env string) bool {
	bval, _ := strconv.ParseBool(env)
	return bval
}

func getFloat(env string) float64 {
	f, _ := strconv.ParseFloat(env, 64)
	return f
}

func getInt(val string) int {
	i, _ := strconv.Atoi(val)
	return i
}
