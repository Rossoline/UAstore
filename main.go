package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {
	fmt.Print("Server started on port:" + portNumber)
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_ = http.ListenAndServe(portNumber, nil)
}
