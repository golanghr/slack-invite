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

package main

import (
	"errors"
	"regexp"

	pb "github.com/golanghr/slack-invite/protos"
	"golang.org/x/net/context"
)

// Stats - HTTP(REST API) and gRPC endpoint designed to be used if we ever choose to go with
// full javascript on the index page. I had no will to implement that from the start but is ongoing idea.
// In case that there are issues with validation we will return error
func (s *Service) Stats(ctx context.Context, in *pb.Request) (*pb.Stats, error) {
	s.Entry.Debugf("Received new stats request: %v", in)

	stats, err := s.Slack.GetStatsPb()

	if err != nil {
		s.Entry.Errorf("Failed to retreive stats from slack: %s", err)
	}

	return stats, err
}

// Invite - HTTP(REST API) and gRPC endpoint designed to invite new member to the team.
// In case that there are issues with validation we will return error
//
// Example Responses
//  {"error":"rpc error: code = 2 desc = \"Failed to invite to team: already_in_team\""}
//  or
//  {"Ok":true}
func (s *Service) Invite(ctx context.Context, in *pb.Request) (*pb.Invite, error) {
	s.Entry.Debugf("Received new invite request: %v", in)

	if len(in.FirstName) < 2 {
		return nil, errors.New("First name must be provided in order to send new invite.")
	}

	if len(in.LastName) < 2 {
		return nil, errors.New("Last name must be provided in order to send new invite.")
	}

	ere := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !ere.MatchString(in.Email) {
		return nil, errors.New("Valid email must be provided in order to send new invite.")
	}

	teamName, _ := s.Options.Get("slack-team-name")

	if err := s.InviteToTeam(teamName.String(), in.FirstName, in.LastName, in.Email); err != nil {
		return nil, err
	}

	return &pb.Invite{Ok: true}, nil
}
