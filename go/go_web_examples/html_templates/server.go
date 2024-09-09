package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type SINDetails struct {
	Success bool
	Reason string
}

func main() {
	tmpl := template.Must(template.ParseFiles("form.html"))
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		fmt.Fprintf(w, "Welcome to the %s page, Guest", r.URL.Path[1:])
	})

	http.ListenAndServe(":80", nil)
}

