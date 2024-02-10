package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/waddellmp/gobnb/pkg/config"
	"github.com/waddellmp/gobnb/pkg/render"
)

func main() {

	//-------------------------------------------------------------------------
	// Initialize app config

	var appConfig = config.AppConfig{}
	flag.StringVar(&appConfig.Port, "p", ":8080", "Server port")
	flag.StringVar(&appConfig.Cache, "m", "static", "Cache mode (static/dynamic)")
	flag.Parse()

	//-------------------------------------------------------------------------
	// Initialize Template Rendering

	render.Initialize(appConfig.Cache)

	//-------------------------------------------------------------------------
	// Register Handlers

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	//-------------------------------------------------------------------------
	// Start server

	fmt.Printf("Starting application on port %s\n", appConfig.Port)
	_ = http.ListenAndServe(appConfig.Port, nil)
}
