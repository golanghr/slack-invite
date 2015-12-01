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

import "strings"

// getRestAPIAddr - Used by the http handlers. As REST API is not on the same
// port we need somehow to figure out exact location before we pull for any data
// over angular.
// @TODO - This looks shitty. Find better, more lightweight approach. Maybe add it
// directly into the server itself.
func getRestAPIAddr(rh string) string {
	gra, _ := opts.Get("grpc-rest-addr")
	gras := strings.Split(gra.String(), ":")
	ah := gra.String()

	if gras[0] == "" {
		rhs := strings.Split(rh, ":")
		ah = strings.Join([]string{rhs[0], gras[1]}, ":")
	}

	return ah
}
