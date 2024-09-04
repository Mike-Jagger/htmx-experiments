package main

import (
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	err := p1.save()
	if err != nil {
		fmt.Println("Error while creating file: ", err)
		return
	}
	page, err := loadPage("TestPage")
	if err != nil {
		fmt.Println("Error while reading file: ", err)
		return
	}
	fmt.Println(page)
}