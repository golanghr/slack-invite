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
	"github.com/Sirupsen/logrus"
	"github.com/golanghr/platform/logging"
	"github.com/golanghr/platform/options"
)

const (
	SSL_CERT_FILE = "/tmp/certs/server.crt"
	SSL_KEY_FILE  = "/tmp/certs/server.key"
)

func init() {

	// We will rewrite service SSL options here as we do not want to expose SSL
	// certificates to the WWW. Additionally, it's the bad practice to have such details
	// available within private repos too.
	//if err := ioutil.WriteFile(SSL_CERT_FILE, []byte(strings.Replace(os.Getenv("SSL_CERT"), "\\n", "\n", -1)), 0755); err != nil {
	//	log.Fatalf("Failed to write SSL cert file: %s", err)
	//}

	//if err := ioutil.WriteFile(SSL_KEY_FILE, []byte(strings.Replace(os.Getenv("SSL_KEY"), "\\n", "\n", -1)), 0755); err != nil {
	//	log.Fatalf("Failed to write SSL key file: %s", err)
	//}

	serviceOptions["grpc-tls-cert"] = SSL_CERT_FILE
	serviceOptions["grpc-tls-key"] = SSL_KEY_FILE

	opts, _ = options.New("memo", serviceOptions)
	serviceName, _ := opts.Get("service-name")
	serviceVersion, _ := opts.Get("service-version")

	log = logging.New(opts)
	logger = log.WithFields(logrus.Fields{
		"service": serviceName.String(),
		"version": serviceVersion.Float(),
	})
}
