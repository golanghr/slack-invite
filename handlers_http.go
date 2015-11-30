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
func IndexHandler(rw http.ResponseWriter, req *http.Request) {
	// Stop here if OPTIONS is request.
	// @TODO Add this as middleware...
	if req.Method == "OPTIONS" {
		return
	}

	lp := path.Join("templates", "layout.html")

	templatePath := "index"
	templateFileType := "html"

	if req.URL.Path != "/" {
		templatePath = req.URL.Path
	}

	// Unset file type if we have favicon request.
	if strings.Contains(req.URL.Path, "favicon") {
		templateFileType = ""
	}

	fp := path.Join("templates", strings.Join([]string{templatePath, templateFileType}, "."))

	// Return a 404 if the template doesn't exist
	info, err := os.Stat(fp)

	if err != nil {
		if os.IsNotExist(err) {
			logger.Errorln(err.Error())
			http.NotFound(rw, req)
			return
		}
	}

	// Return a 404 if the request is for a directory
	if info.IsDir() {
		http.NotFound(rw, req)
		return
	}

	tmpl, err := template.New("slackinvite").Delims("[[", "]]").ParseFiles(lp, fp)

	if err != nil {
		logger.Errorln(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(rw, http.StatusText(500), 500)
		return
	}

	if err := tmpl.ExecuteTemplate(rw, "layout", nil); err != nil {
		logger.Errorln(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(rw, http.StatusText(500), 500)
	}
}
