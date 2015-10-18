// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package main ...
package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// IndexHandler -
func IndexHandler(w http.ResponseWriter, req *http.Request) {

	indext, err := template.ParseFiles(
		"./templates/index.html", "./templates/header.html", "./templates/footer.html",
	)

	errort, err := template.ParseFiles(
		"./templates/error.html", "./templates/header.html", "./templates/footer.html",
	)

	if err != nil {
		errort.Execute(w, map[string]interface{}{
			"Title": fmt.Sprintf("Error Occurred | %s", slackinvite.Name()),
		})
		return
	}

	indext.Execute(w, map[string]interface{}{
		"Title": fmt.Sprintf("Request Slack Invitation | %s", slackinvite.Name()),
	})
}
