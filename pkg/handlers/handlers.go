package handlers

import (
	"github.com/onlydarel/sea-salon/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "index.page.html")
}

// Assets is a handler for static asset files in the templates folder
func Assets(w http.ResponseWriter, r *http.Request) {
	// Create the handler for serving static files
	fileServer := http.FileServer(http.Dir("templates/assets"))
	http.StripPrefix("/assets/", fileServer).ServeHTTP(w, r)
}

// Styles is a handler for css files in the templates folder
func Styles(w http.ResponseWriter, r *http.Request) {
	fileServer := http.FileServer(http.Dir("templates/styles"))
	http.StripPrefix("/styles/", fileServer).ServeHTTP(w, r)
}
