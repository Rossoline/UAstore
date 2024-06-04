package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandlerFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello world!")
	})
}
