package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/waddellmp/gobnb/pkg/config"
	"github.com/waddellmp/gobnb/pkg/handlers"
	"github.com/waddellmp/gobnb/pkg/render"
)

const portNumber = ":8080"

func main() {

	//=========================================================================

	// initialize app config

	var app config.AppConfig = config.AppConfig{
		TemplateCache: make(map[string]*template.Template),
	}

	//=========================================================================

	// build static cache

	c, err := render.BuildStaticCache()
	if err != nil {
		log.Fatalln("unable to create template cache", err)
	}

	app.TemplateCache = c

	//=========================================================================

	// setup handlers

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s\n", portNumber)

	//=========================================================================

	// start server

	_ = http.ListenAndServe(portNumber, nil)
}
