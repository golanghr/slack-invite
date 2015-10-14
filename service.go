// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package main ...
package main

import (
	"github.com/golanghr/platform/config"
	"github.com/golanghr/platform/service"
)

// SlackInvite -
type SlackInvite struct {
	config.Manager
	service.Service
}

// LoadConfiguration -
func (si *SlackInvite) LoadConfiguration(cnf map[string]interface{}) error {
	log.Debug("Starting configuration management ...")

	if si.Manager, err = config.New(cnf); err != nil {
		return err
	}

	// Setup defaults ...

	if err := si.Manager.EnsureSet("service-name", ServiceName); err != nil {
		return err
	}

	if err := si.Manager.EnsureSet("service-description", ServiceDescription); err != nil {
		return err
	}

	if err := si.Manager.EnsureSet("service-version", ServiceVersion); err != nil {
		return err
	}

	return nil
}

// LoadService -
func (si *SlackInvite) LoadService() (err error) {
	log.Debug("Starting service management ...")

	if si.Service, err = service.New(si.Manager); err != nil {
		return
	}

	return
}

// Run -
func (si *SlackInvite) Run() (err error) {
	log.Warn("Starting service ...")

	return
}

// Recover -
func (si *SlackInvite) Recover() {
	if err := recover(); err != nil {
		errlog.Fatal("Panic happen! Killing service now...")
	}
}

// NewSlackInvite -
func NewSlackInvite(cnf map[string]interface{}) (*SlackInvite, error) {
	slackinvite := new(SlackInvite)

	if err := slackinvite.LoadConfiguration(cnf); err != nil {
		return nil, err
	}

	if err := slackinvite.LoadService(); err != nil {
		return nil, err
	}

	return slackinvite, nil
}
