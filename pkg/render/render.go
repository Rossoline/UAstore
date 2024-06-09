package render

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl + ".page.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("eror parsing template:", err)
	}
}
