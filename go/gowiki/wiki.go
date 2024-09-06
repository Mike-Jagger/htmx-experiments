package main

import (
	"fmt"
	"gowiki/handlers"
	"gowiki/shared"
	"log"
	"net/http"
)

func MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := shared.ValidPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}

		fn(w, r, m[2])
	}
}

func main() {
	http.HandleFunc("/", handlers.Handler)
	http.HandleFunc("/view/", MakeHandler(handlers.ViewHandler))
	http.HandleFunc("/edit/", MakeHandler(handlers.EditHandler))
	http.HandleFunc("/save/", MakeHandler(handlers.SaveHandler))

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