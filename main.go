package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is home page!")
}

// About is about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValue(5, 7)
	fmt.Fprintf(w, "This is about page and sum is - %d!", sum)
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(20.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "Cant divide by 0")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f devided by %f is %f", 100.0, 0.0, f))
}

func divideValues(x, y float32) (float32, error) {
	if x <= 0 || y <= 0 {
		er := errors.New("cant divide by zero")
		return 0, er
	}
	result := x / y
	return result, nil
}

func addValue(x, v int) int {
	sum := x + v
	return sum
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)
	_ = http.ListenAndServe(portNumber, nil)
}
