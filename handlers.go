package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home")
}

// About is about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl + ".page.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("eror parsing template:", err)
	}
}
