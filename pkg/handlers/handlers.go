package handlers

import (
	"net/http"
	"uastore/pkg/render"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home")
}

// About is about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about")
}
