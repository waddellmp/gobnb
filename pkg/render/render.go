package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Template cache
var tc = make(map[string]*template.Template)

// Render Template with caching
func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// Check if template is in cache
	_, isCached := tc[t]

	// If template is not in cache, add it
	// Handle error if template fails to be added to cache
	if !isCached {
		log.Printf("adding teplate %s to cache.\n", t)
		err = addTemplateToCache(t)

		if err != nil {
			log.Println(err.Error())

			// Stop here
			return
		}
	} else {
		fmt.Printf("template %s is present in cache", t)
	}

	// Get template from cache
	tmpl = tc[t]

	// Write parsed template to response writer
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
	}
}

// Add template to cache, returns error if template fails to parse
// Usage: addTemplateToCache("home.page.html")
func addTemplateToCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base-layout.html",
	}

	fmt.Printf("templates: %v\n", templates)

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	tc[t] = tmpl
	return nil
}
