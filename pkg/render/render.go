package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// pkg configuration
type renderCache struct {
	cache map[string]*template.Template
}

var rc = renderCache{}

// Builds the render cache and returns it
func BuildStaticCache() (map[string]*template.Template, error) {
	// create a cache
	c := make(map[string]*template.Template)

	// set the cache
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

	// ========================================================================
	// try to get the templace from cache
	t, ok := rc.cache[filename]
	if !ok {
		log.Fatalln("template is not in the cache for some reason")
	}

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
