package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, filename string) {
	// create a template cache
	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatalln("error creating template cache", err)
	}

	// try to get the template from cache
	t, ok := tc[filename]
	if !ok {
		log.Fatalln("template is not in the cache for some reason")
	}

	// create new buffer
	buf := new(bytes.Buffer)

	// use empty buffer to execute template and test if error
	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	// no error present so execute against response writer
	t.Execute(w, nil)
}

func CreateTemplateCache() (map[string]*template.Template, error) {

	// construct map of filesNames -> *template.Template
	cache := map[string]*template.Template{}

	// get slice of filenames (ie: home-page.html) that were found in filesystem
	pages, err := filepath.Glob("./templates/*-page.html")
	if err != nil {
		return cache, err
	}

	// Loop through page paths
	for _, page := range pages {

		// Get name of the file from matches (ie: templates/home-page.html)
		name := filepath.Base(page)

		// Construct a new template with name and parse template file path
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		// Get layout filenames
		layoutFiles, err := filepath.Glob("./templates/*-layout.html")
		if err != nil {
			return cache, err
		}

		// Parse in the layout file(s) file using ParseGlob()
		if len(layoutFiles) > 0 {
			ts, err = ts.ParseGlob("./templates/*-layout.html")
			if err != nil {
				return cache, err
			}
		}

		// Insert template in cache
		cache[name] = ts
	}

	return cache, nil
}
