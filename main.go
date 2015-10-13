// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package main ...
package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/golanghr/platform/config"
	"github.com/golanghr/platform/logging"
	"github.com/golanghr/platform/server"
	"github.com/golanghr/platform/service"
)

var (
	log          logging.Logging
	conf         config.Manager
	err          error
	serv         service.Service
	http         server.Server
	slog         *logrus.Entry
	elog         *logrus.Entry
	serviceEvent chan string
)

func main() {
	defer recovery()

	slog.Debug("Hello! We are going to prepare service now ...")

	if conf, err = config.New(serviceConfig); err != nil {
		elog.Fatal("Configuration error happen! Killing service now...")
	}

	// Just small helper located bellow main() to keep main looking ok
	setupConfig()

	if serv, err = service.New(conf); err != nil {
		elog.Fatal("Service initialization error happen! Killing service now...")
	}

	if http, err = server.NewHTTPServer(serv); err != nil {
		elog.Fatal("HTTP server initialization error happen! Killing service now...")
	}

	serviceEvent = make(chan string)
	go handleServiceEventMessages(serviceEvent)

	if err := http.Start(serviceEvent); err != nil {
		elog.Fatal("HTTP server startup error happen! Killing service now...")
	}
}

func handleServiceEventMessages(se chan string) {
	for {
		event := <-se

		if event == server.STARTING {
			slog.Warning("Starting up HTTP/HTTPS service now ...")
		}

	}
}

func setupConfig() {
	if _, err := conf.GetOrSet("service-name", ServiceName); err != nil {
		elog.Fatal("Configuration error happen! Killing service now...")
	}

	if _, err := conf.GetOrSet("service-description", ServiceName); err != nil {
		elog.Fatal("Configuration error happen! Killing service now...")
	}

	if _, err := conf.GetOrSet("service-version", ServiceName); err != nil {
		elog.Fatal("Configuration error happen! Killing service now...")
	}

	if _, err := conf.GetOrSet("server-http-tls", "yes"); err != nil {
		elog.Fatal("Configuration error happen! Killing service now...")
	}
}

func recovery() {
	err := recover()

	if err != nil {
		log.WithFields(logrus.Fields{
			"service": ServiceName,
			"err":     err,
		}).Fatal("Panic happen! Killing service now...")
	}
}

// init - used to setup logger
func init() {
	log = logging.New(map[string]interface{}{
		"formatter": "text",
		"level":     logrus.DebugLevel,
	})

	slog = log.WithFields(logrus.Fields{"service": ServiceName})
	elog = log.WithFields(logrus.Fields{"service": ServiceName, "err": err})
}
