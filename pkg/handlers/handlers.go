package handlers

import (
	"net/http"

	"github.com/waddellmp/gobnb/pkg/render" // Import the package that contains the RenderTemplate function
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home-page.html") // Use the RenderTemplate function from the imported package
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about-page.html") // Use the RenderTemplate function from the imported package
}
