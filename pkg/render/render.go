package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type renderCache struct {
	cache map[string]*template.Template
	mode  string
}

var rc = renderCache{}

func Initialize(mode string) {
	rc.mode = mode
	if rc.mode == "static" {
		fmt.Println("Building static cache")
		BuildCache()
		return
	}
}

// builds the render cache and returns it
func BuildCache() (map[string]*template.Template, error) {
	// make cache and set cache
	c := make(map[string]*template.Template)
	rc.cache = c

	// get slice of filenames (ie: home-page.html) that were found in filesystem
	pages, err := filepath.Glob("./templates/*-page.html")
	if err != nil {
		return rc.cache, err
	}

	// loop through page paths
	for _, page := range pages {

		// get name of the file from matches (ie: templates/home-page.html)
		name := filepath.Base(page)

		// construct a new template with name and parse template file path
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return rc.cache, err
		}

		// get layout filenames
		layoutFiles, err := filepath.Glob("./templates/*-layout.html")
		if err != nil {
			return rc.cache, err
		}

		// parse in the layout file(s) file using ParseGlob()
		if len(layoutFiles) > 0 {
			ts, err = ts.ParseGlob("./templates/*-layout.html")
			if err != nil {
				return rc.cache, err
			}
		}

		// insert template in cache
		rc.cache[name] = ts
	}

	return rc.cache, nil
}

func RenderTemplate(w http.ResponseWriter, filename string) {
	// Build dynamic cache on every request (development mode)
	if rc.mode != "static" {
		fmt.Println("Dynamically building cache")
		BuildCache()
	}
	t, ok := rc.cache[filename]
	if !ok {
		log.Fatalln("Unable to find template")
	}
	// --------------------------------------------------------------------
	// try to get the template from cache

	// create new buffer
	buf := new(bytes.Buffer)

	// use empty buffer to execute template and test if error
	err := t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	// no error present so execute against response writer
	t.Execute(w, nil)
}
