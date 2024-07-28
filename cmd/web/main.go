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
	fmt.Print("Server started on port:" + portNumber)
	http.HandleFunc("/", handlers.Home)
	fmt.Println("Hello")
	http.HandleFunc("/about", handlers.About)
	_ = http.ListenAndServe(portNumber, nil)
}
