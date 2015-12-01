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

// Slack - Just small wrapper on top of existing slack api to help us with protobuff
// rendering. Additionally, it's here if we need to implement more out of slack api that is not
// available under existing slack.Client
type Slack struct {
	*slack.Client

	Token string
	Debug bool
}

// GetStatsPb - Will return back stats protobuff that is necessary for both grpc,
// REST API and HTML templating
func (s *Slack) GetStatsPb() (*pb.Stats, error) {
	users, err := s.GetUsers("1")

	if err != nil {
		return nil, err
	}

	stats := &pb.Stats{
		Active: []string{},
		Away:   []string{},
		Admins: []string{},
		Total:  int64(len(users)),
	}

	for _, user := range users {
		if user.Presence == "active" && user.RealName != "" {
			stats.Active = append(stats.Active, user.RealName)
		} else if user.Presence == "away" && user.RealName != "" {
			stats.Away = append(stats.Away, user.RealName)
		}

		if user.IsAdmin {
			stats.Admins = append(stats.Admins, user.RealName)
		}
	}

	return stats, nil
}

// NewSlack - Will return back new Slack that will be later on used by the Service{}
func NewSlack(token string, debug bool) *Slack {
	api := slack.New(token)
	api.SetDebug(debug)

	return &Slack{
		Client: api,
		Token:  token,
		Debug:  debug,
	}
}
