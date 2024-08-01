package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is home page!")
	fmt.Fprintf(w, "This is home page!")
}

// About is about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValue(5, 7)
	fmt.Fprintf(w, "This is about page and sum is - %d!", sum)
	fmt.Fprintf(w, "This is about page and sum is - %d!", sum)
}

func addValue(x, v int) int {
	sum := x + v
	return sum
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_ = http.ListenAndServe(portNumber, nil)
}
