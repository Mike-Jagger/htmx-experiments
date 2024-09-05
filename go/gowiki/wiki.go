package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var (
	templates = template.Must(template.ParseFiles("templ/edit.html", "templ/view.html"))
	validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
)

type Page struct {
	Title string
	Body []byte
	Markup template.HTML
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

// func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
// 	m := validPath.FindStringSubmatch(r.URL.Path)
// 	if m == nil {
// 		http.NotFound(w, r)
// 		return "", errors.New("Invalid Page Title")
// 	}
// 	return m[2], nil
// }

func substitute(b []byte) []byte {
	pageName := string(b)
	return []byte("<a href='/view/"+pageName+"'>"+pageName+"</a>")
}

func loadPage(title string) (*Page, error) {
	filename := "data/" + title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	files, err := os.ReadDir("data")
	if err != nil {
        return nil, err
    }

    var fileNames []string
    for _, file := range files {
        if file.IsDir() {
            continue
        }
		fileName := strings.TrimSuffix(file.Name(), ".txt")
        fileNames = append(fileNames, fileName)
    }
	
	pattern := fmt.Sprintf("\\b(%s)\\b", strings.Join(fileNames, "|"))
	body = regexp.MustCompile(pattern).ReplaceAllFunc(body, substitute)
	
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	p.Markup = template.HTML(p.Body)
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		fmt.Println("Error parsing", tmpl, ".html file:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}

		fn(w, r, m[2])
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		fmt.Println("Can't find page:", title)
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
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
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

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