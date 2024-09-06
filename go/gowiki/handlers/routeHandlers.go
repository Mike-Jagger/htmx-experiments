package handlers

import (
	"fmt"
	"gowiki/shared"
	"gowiki/util"
	"net/http"
)

type Page shared.Page

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func ViewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := util.LoadPage(title)
	if err != nil {
		fmt.Println("Can't find page:", title, err)
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	util.RenderTemplate(w, "view", p)
}

func EditHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := util.LoadPage(title)
	if err != nil {
		fmt.Println("Can't find page:", title, err)
		p = &util.Page{Title: title}
	}

	util.RenderTemplate(w, "edit", p)
}

func SaveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &shared.Page{Title: title, Body: []byte(body)}

	err := p.Save()

	if err != nil {
		fmt.Println("Error saving file:", title, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
