package util

import (
	"fmt"
	"gowiki/shared"
	"html/template"
	"net/http"
	"os"
	"regexp"
	"strings"
)

// func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
// 	m := validPath.FindStringSubmatch(r.URL.Path)
// 	if m == nil {
// 		http.NotFound(w, r)
// 		return "", errors.New("Invalid Page Title")
// 	}
// 	return m[2], nil
// }

type Page shared.Page

func substitute(b []byte) []byte {
	pageName := string(b)
	return []byte("<a href='/view/"+pageName+"'>"+pageName+"</a>")
}

func LoadPage(title string) (*Page, error) {
	filename := "./../data/" + title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	files, err := os.ReadDir("./../data")
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

func RenderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	p.Markup = template.HTML(p.Body)
	err := shared.Templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		fmt.Println("Error parsing", tmpl, ".html file:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}