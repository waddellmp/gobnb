package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/waddellmp/gobnb/pkg/config"
	"github.com/waddellmp/gobnb/pkg/render"
)

func main() {

	//=========================================================================

	// initialize app config

	var appConfig config.AppConfig

	port := flag.String("p", ":8080", "Server port")
	flag.Parse()

	appConfig.Port = *port

	//=========================================================================

	// build static cache

	_, err := render.BuildStaticCache()
	if err != nil {
		log.Fatalln("unable to create template cache", err)
	}

	//=========================================================================

	// setup handlers

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting application on port %s\n", *port)

	//=========================================================================

	// start server

	_ = http.ListenAndServe(*port, nil)
}
