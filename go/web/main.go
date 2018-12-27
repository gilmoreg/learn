package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hi</h1>")
}

func main() {
	fmt.Println("Starting server on port 3000")
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
