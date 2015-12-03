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
	"net/http"
	"os"
	"path"
	"strings"
	"text/template"
)

// IndexHandler -
func IndexHandler(rw http.ResponseWriter, req *http.Request) {
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

	gaua, _ := opts.Get("google-analytics-tracking-id")

	params := map[string]interface{}{
		"title":   "Golang.hr automated slack invitation as easy as 1,2,3!",
		"apihost": "ghrslack-rest.ngrok.com",
		"gaua":    gaua.String(),
	}

	if err := tmpl.ExecuteTemplate(rw, "layout", params); err != nil {
		logger.Errorln(err.Error())
		http.Error(rw, http.StatusText(500), 500)
	}
}
