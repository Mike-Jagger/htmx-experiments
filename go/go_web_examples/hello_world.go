package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Guest! You requested this resource: %s", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)

	http.ListenAndServe(":80", nil)
	fmt.Println("Listening on port :80")
}