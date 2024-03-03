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
	cache         map[string]*template.Template
	mode          string
	pagePattern   string
	layoutPattern string
}

var rc renderCache = renderCache{
	cache: make(map[string]*template.Template),
}

func SetMode(mode string) {
	rc.mode = mode
}
func SetLayouts(layout string) {
	rc.layoutPattern = layout
}
func SetPages(page string) {
	rc.pagePattern = page
}

// builds the render cache and returns it
func BuildStaticCache() (map[string]*template.Template, error) {
	// get all template pages filenames
	pages, err := filepath.Glob(rc.pagePattern)
	if err != nil {
		log.Printf("Error getting pages: %v\n", err)
		return rc.cache, err
	}

	// get all template layout filenames
	layouts, err := filepath.Glob(rc.layoutPattern)
	if err != nil {
		log.Println("Error getting pages: ", err)
		return rc.cache, err
	}

	// loop through all pages, parse each with layout files, and cache
	for _, page := range pages {
		pageName := filepath.Base(page)

		// check if page is already in cache
		_, ok := rc.cache[pageName]

		// add to cache
		if !ok {
			// parse file
			pt, err := template.ParseFiles(page)
			if err != nil {
				log.Println("Error parsing page templates: ", err)
				return rc.cache, err
			}

			// parse page template with layout files
			_, err = pt.ParseFiles(layouts...)
			if err != nil {
				log.Println("Error parsing layout templates: ", err)
				return rc.cache, err
			}

			// cache the template
			rc.cache[pageName] = pt
		}
	}

	return rc.cache, nil
}

func RenderTemplate(w http.ResponseWriter, filename string) {
	t, ok := rc.cache[filename]

	if !ok {
		log.Println("No file found")
		return
	}

	fmt.Printf("Template: %v\n", t)
	// create new buffer
	buf := new(bytes.Buffer)

	// use empty buffer to execute template and test if error
	execErr := t.Execute(buf, nil)

	if execErr != nil {
		log.Println(execErr)
	}

	// no error present so execute against response writer
	t.Execute(w, nil)
}
