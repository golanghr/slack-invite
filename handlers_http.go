// Copyright 2015 The Golang.hr Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package main ...
package main

import (
	"net/http"
	"os"
	"path"
	"strings"
	"text/template"
)

// IndexHandler -
func IndexHandler(w http.ResponseWriter, req *http.Request) {
	lp := path.Join("templates", "layout.html")

	templatePath := "index"

	if req.URL.Path != "/" {
		templatePath = req.URL.Path
	}

	fp := path.Join("templates", strings.Join([]string{templatePath, "html"}, "."))

	// Return a 404 if the template doesn't exist
	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			// Log the detailed error
			logger.Errorln(err.Error())
			http.NotFound(w, req)
			return
		}
	}

	// Return a 404 if the request is for a directory
	if info.IsDir() {
		http.NotFound(w, req)
		return
	}

	tmpl, err := template.ParseFiles(lp, fp)

	if err != nil {
		// Log the detailed error
		logger.Println(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "layout", nil); err != nil {
		logger.Errorln(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}
