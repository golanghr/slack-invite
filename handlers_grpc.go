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
	pb "github.com/golanghr/slack-invite/protos"
	"golang.org/x/net/context"
)

// TeamDetails - Depending on server type (http, grpc, ...) will return back new simple response.
// If this is HTTP, it will return back application/json with HelloWorld being serialized into json
// If this is GRPC, it will return back entire hello.HelloWorld as bytes
func (s *Service) TeamDetails(ctx context.Context, in *pb.SlackInviteRequest) (*pb.SlackInvite, error) {
	return &pb.SlackInvite{}, nil
}
