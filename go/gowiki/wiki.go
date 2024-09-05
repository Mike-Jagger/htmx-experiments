package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var (
	templates = template.Must(template.ParseFiles("edit.html", "view.html"))
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

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.Execute(w, tmpl+".html", p)
	if err != nil {
		fmt.Println("Error parsing", tmpl, ".html file:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		fmt.Println("Can't find page:", title)
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}

	err := p.save()

	if err != nil {
		fmt.Println("Error saving file:", title)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)

	fmt.Println("Server starting on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	// p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	// err := p1.save()
	// if err != nil {
	// 	fmt.Println("Error while creating file: ", err)
	// 	return
	// }
	// page, err := loadPage("TestPage")
	// if err != nil {
	// 	fmt.Println("Error while reading file: ", err)
	// 	return
	// }
	// fmt.Println(page)

}