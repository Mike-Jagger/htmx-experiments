package shared

import (
	"html/template"
	"os"
)

type Page struct {
	Title  string
	Body   []byte
	Markup template.HTML
}

func (p *Page) Save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}