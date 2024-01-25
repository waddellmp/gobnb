package main

import (
	"fmt"
	"net/http"

	"github.com/waddellmp/gobnb/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s\n", portNumber)
	// Listen on port and handle requests & use default serve mux by passing nil handler
	_ = http.ListenAndServe(portNumber, nil)
}
