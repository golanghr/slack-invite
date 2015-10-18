// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package main ...
package main

import (
	"net/http"
	"time"

	"github.com/golanghr/platform/config"
	"github.com/golanghr/platform/service"
	"github.com/gorilla/mux"
)

// SlackInvite -
type SlackInvite struct {
	service.Service

	Config config.Manager

	Server  *http.Server
	Handler *mux.Router

	Quit chan bool
}

// AddHandler -
func (si *SlackInvite) AddHandler(handler *mux.Router) {
	si.Handler = handler
}

// LoadConfiguration - Will basically ensure that service defaults are set and that
// they are properly syncronized against etcd
func (si *SlackInvite) LoadConfiguration(cnf map[string]interface{}) error {
	log.Debug("Ensuring configuration ...")

	cm := map[string]string{
		"service-name":        ServiceName,
		"service-description": ServiceDescription,
		"service-version":     ServiceVersion,
		"addr":                ServiceAddr,
		"read_timeout":        ServiceReadTimeout,
		"write_timeout":       ServiceWriteTimeout,
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

// Run -
func (si *SlackInvite) Run() (err error) {
	var addr string
	var rTimeout time.Duration
	var wTimeout time.Duration

	if addr, err = si.Config.GetString("addr"); err != nil {
		return
	}

	if rTimeout, err = si.Config.GetDuration("read_timeout"); err != nil {
		return
	}

	if wTimeout, err = si.Config.GetDuration("write_timeout"); err != nil {
		return
	}

	log.Warnf("Starting HTTP service on (addr: %s) ...", addr)

	si.Server = &http.Server{
		Addr:         addr,
		Handler:      si.Handler,
		ReadTimeout:  rTimeout * time.Second,
		WriteTimeout: wTimeout * time.Second,
	}

	return si.Server.ListenAndServe()
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

	return slackinvite, nil
}
