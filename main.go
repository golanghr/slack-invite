// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package main ...
package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/golanghr/platform/config"
	"github.com/golanghr/platform/logging"
	"github.com/golanghr/platform/service"
)

var (
	log  logging.Logging
	conf config.Manager
	err  error
	serv service.Service
)

func init() {
	log = logging.New(map[string]interface{}{
		"formatter": "text",
		"level":     logrus.DebugLevel,
	})
}

func main() {
	slog := log.WithFields(logrus.Fields{"service": "slack-invite"})
	defer recovery()

	slog.Debug("Hello! We are going to prepare service now ...")

	if conf, err = config.New(serviceConfig); err != nil {
		log.WithFields(logrus.Fields{
			"service": "slack-invite",
			"err":     err,
		}).Fatal("Configuration error happen! Killing service now...")
	}

}

func recovery() {
	err := recover()

	if err != nil {
		log.WithFields(logrus.Fields{
			"service": "slack-invite",
			"err":     err,
		}).Fatal("Panic happen! Killing service now...")

		// Exit under critical conditions
		os.Exit(2)
	}
}
