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
	"github.com/golanghr/slack"
	pb "github.com/golanghr/slack-invite/protos"
)

/**
type Team struct {
	Ok bool
}
**/

// Slack - Just small wrapper on top of existing slack api.
type Slack struct {
	*slack.Client

	Token string
	Debug bool

	Users []slack.User
}

// GetSlackInvitePb -
func (s *Slack) GetSlackInvitePb() (*pb.SlackInvite, error) {
	var err error

	if s.Users, err = s.GetUsers("1"); err != nil {
		return nil, err
	}

	active := []string{}
	away := []string{}
	admins := []string{}

	for _, user := range s.Users {
		if user.Presence == "active" {
			active = append(active, user.RealName)
		}

		if user.Presence == "away" {
			away = append(away, user.RealName)
		}

		if user.IsAdmin {
			admins = append(admins, user.RealName)
		}
	}

	return &pb.SlackInvite{
		Active: active,
		Away:   away,
		Admins: admins,
		Total:  int64(len(s.Users)),
	}, nil
}

// NewSlack -
func NewSlack(token string, debug bool) *Slack {

	api := slack.New(token)
	api.SetDebug(debug)

	return &Slack{
		Client: api,
		Token:  token,
		Debug:  debug,
	}
}
