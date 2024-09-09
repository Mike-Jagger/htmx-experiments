package main

import (
	"html/template"
	"net/http"
)

type SINDetails struct {
	Success bool
	SIN string
	Reason string
}

func main() {
	tmpl := template.Must(template.ParseFiles("form.html"))
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		tmpl.Execute(w, SINDetails{Success: true, SIN: r.FormValue("SIN"), Reason: r.FormValue("reason")})
	})

	http.ListenAndServe(":80", nil)
}

