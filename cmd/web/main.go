package main

import (
	"fmt"
	"net/http"

	"github.com/Rossoline/UAstore/cmd/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	fmt.Print("Server started on port:" + portNumber)
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	_ = http.ListenAndServe(portNumber, nil)
}
