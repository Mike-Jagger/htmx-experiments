package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "Authorization")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	if likesCake, ok := session.Values["likesCake"].(bool); !ok{
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	} else if likesCake {
		fmt.Fprintf(w, "You like cake, and it is definitely not a lie")
		return
	}

	fmt.Fprintf(w, "The cake is a lie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "Authorization")

	session.Values["authenticated"] = true
	session.Values["likesCake"] = true
	session.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "Authorization")

	session.Values["authenticated"] = false
	session.Save(r, w)
}

type User struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Age int `json:"age"`
}

func main() {
	http.HandleFunc("/secret", secret)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		peter := User{
			FirstName: "peter",
			LastName: "parker",
			Age: 25,
		}

		json.NewEncoder(w).Encode(peter)
	})
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var user User
		json.NewDecoder(r.Body).Decode(&user)

		fmt.Fprintf(w, "%s %s is %d years old", user.FirstName, user.LastName, user.Age)
	})
	http.ListenAndServe(":80", nil)
}