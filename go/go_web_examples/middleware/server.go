package main

import (
	"fmt"
	"log"
	"net/http"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("bar")
}

func bazz(w http.ResponseWriter, r *http.Request) {
	fmt.Println("bazz")
} 


func main() {
	http.HandleFunc("/foo", logging(foo))
	http.HandleFunc("/bar", logging(bar))
	http.HandleFunc("/bazz", logging(bazz))

	http.ListenAndServe(":80", nil)
}	