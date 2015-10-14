// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package main ...
package main

import (
	"runtime"

	"github.com/Sirupsen/logrus"
	"github.com/golanghr/platform/logging"
	"github.com/golanghr/platform/utils"
)

var (
	err error

	logger logging.Logging
	log    *logrus.Entry
	errlog *logrus.Entry

	slackinvite *SlackInvite
)

func init() {
	logger = logging.New(map[string]interface{}{
		"formatter": "text",
		"level":     logrus.DebugLevel,
	})

	log = logger.WithFields(logrus.Fields{"service": ServiceName})
	errlog = logger.WithFields(logrus.Fields{"service": ServiceName, "err": err})
}

func main() {
	log.Debug("Hello! We are going to nuke service now ...")

	runtime.GOMAXPROCS(utils.GetProcessCount("SLACK_INVITE_PROCESS_COUNT"))

	if slackinvite, err = NewSlackInvite(serviceCnf); err != nil {
		errlog.Fatal("Service establishment error occurred. Terminating service now...")
	}

	defer slackinvite.Recover()

	if err = slackinvite.Run(); err != nil {
		errlog.Fatal("Service runtime error occurred. Terminating service now...")
	}

}
