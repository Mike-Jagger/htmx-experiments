package shared

import (
	"regexp"
	"text/template"
)

var (
	Templates = template.Must(template.ParseFiles("./../templ/edit.html", "./../templ/view.html"))
	ValidPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
)