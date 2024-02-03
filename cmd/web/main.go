package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/waddellmp/gobnb/pkg/config"
	"github.com/waddellmp/gobnb/pkg/render"
)

// test
func main() {

	//-------------------------------------------------------------------------
	// Initialize app config

	var appConfig config.AppConfig

	port := flag.String("p", ":8080", "Server port")
	flag.Parse()

	appConfig.Port = *port

	//-------------------------------------------------------------------------
	// Build Static Cache

	_, err := render.BuildStaticCache()
	if err != nil {
		log.Fatalln("unable to create template cache", err)
	}

	//-------------------------------------------------------------------------
	// Register Handlers

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting application on port %s\n", *port)

	//=========================================================================
	// Start server

	_ = http.ListenAndServe(*port, nil)
}
