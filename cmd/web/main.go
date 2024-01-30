package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/waddellmp/gobnb/pkg/config"
	"github.com/waddellmp/gobnb/pkg/handlers"
	"github.com/waddellmp/gobnb/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	//=========================================================================
	// create static template cache

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalln("unable to create template cache", err)
	}

	//=========================================================================
	// assign template cache to AppConfig
	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s\n", portNumber)
	// Listen on port and handle requests & use default serve mux by passing nil handler
	_ = http.ListenAndServe(portNumber, nil)
}
