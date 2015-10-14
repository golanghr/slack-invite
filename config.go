// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package main ...
package main

import (
	"time"

	etcd "github.com/coreos/etcd/client"
	"github.com/golanghr/platform/utils"
)

var (
	serviceCnf = map[string]interface{}{
		"env":                utils.GetFromEnvOr("SERVICE_SLACK_INVITE_ENV", "sandbox"),
		"folder":             utils.GetFromEnvOr("SERVICE_SLACK_INVITE_ETCD_FOLDER", "golanghr/slack-invite"),
		"auto_sync":          true,
		"auto_sync_interval": 10 * time.Second,
		"etcd": map[string]interface{}{
			"version": "v2",
			"endpoints": []string{
				utils.GetFromEnvOr("SERVICE_SLACK_INVITE_ETCD_ENDPOINT_PRIMARY", "http://localhost:2379"),
				utils.GetFromEnvOr("SERVICE_SLACK_INVITE_ETCD_ENDPOINT_SECONDARY", "http://localhost:2379"),
			},
			"transport":                  etcd.DefaultTransport,
			"username":                   utils.GetFromEnvOr("SERVICE_SLACK_INVITE_ETCD_USERNAME", ""),
			"password":                   utils.GetFromEnvOr("SERVICE_SLACK_INVITE_ETCD_PASSWORD", ""),
			"header_timeout_per_request": time.Second,
		},
	}
)
