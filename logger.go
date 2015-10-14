// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package main ...
package main

import "github.com/Sirupsen/logrus"

// LogFatalError -
func LogFatalError(err error, additional string) {
	logger.WithFields(logrus.Fields{"service": ServiceName, "err": err}).Fatal(additional)
}
