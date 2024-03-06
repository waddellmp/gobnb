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
	// Parse command line flags into AppConfig

	flag.BoolVar(&config.AppConfig.UseCache, "cache", false, "Use template cache")
	flag.StringVar(&config.AppConfig.Port, "p", ":8080", "Server port")
	flag.Parse()

	//-------------------------------------------------------------------------
	// Set pages, layouts, and build static cache for templates on startup

	render.SetPages("./templates/*-page.html")
	render.SetLayouts("./templates/*-layout.html")
	render.BuildStaticCache(false)

	//-------------------------------------------------------------------------
	// Set use cache

	render.SetUseCache(config.AppConfig.UseCache)

	//-------------------------------------------------------------------------
	// Register Handlers

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	//-------------------------------------------------------------------------
	// Start server

	fmt.Printf("Starting application on port %s\n", config.AppConfig.Port)
	_ = http.ListenAndServe(config.AppConfig.Port, nil)
}
