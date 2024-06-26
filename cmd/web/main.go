package main

import (
	"fmt"
	"log"
	"net/http"
	"uastore/pkg/config"
	"uastore/pkg/handlers"
	"uastore/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	fmt.Print("Server started on port:" + portNumber)
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	_ = http.ListenAndServe(portNumber, nil)
	fmt.Println("All good")
}
