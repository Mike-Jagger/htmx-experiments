package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Fprintf(w, "You are looking at page %s of book %s", vars["page"], vars["title"])
	})

	r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
	r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
	r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

	r.HandleFunc("/books/{title}", TestBookHandler).Host("staging.my-website.com")

	r.HandleFunc("/secure", SecureHandler).Schemes("https")
	r.HandleFunc("/insecure", InsecureHandler).Schemes("http")

	specialBookRouter := r.PathPrefix("/special-books").Subrouter().Schemes("https").Host("admin.my-website.com")
	specialBookRouter.HandleFunc("/", AllBooks)
	specialBookRouter.HandleFunc("/{title}", GetBook)

	http.ListenAndServe(":80", r)
	fmt.Println("Listening on port :80")
}