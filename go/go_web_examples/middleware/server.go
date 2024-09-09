package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

type MiddleWare func(http.HandlerFunc) http.HandlerFunc
func Notify(logger *log.Logger, message string) MiddleWare {
	return func(h http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Println("Starting:", message)
			defer logger.Println("Done:", message)
			h(w, r)
		})
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Foo")

}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("bar")
}

func bazz(w http.ResponseWriter, r *http.Request) {
	fmt.Println("bazz")
} 

func handle(f http.HandlerFunc, middlewares ...MiddleWare) http.HandlerFunc {
	for i := range middlewares {
		f = middlewares[len(middlewares)-i-1](f)
		// f = m(f) this will execute the functions in reverse order (stack) before executing first handlerfunc
	}
	return f

}


func main() {
	logger := log.New(os.Stdout, "server: ", log.Lshortfile)

	http.HandleFunc("/foo", logging(foo))
	http.HandleFunc("/bar", logging(bar))
	http.HandleFunc("/bazz", logging(bazz))

	http.Handle("/", Notify(logger, "1")(Notify(logger, "2")(foo)))
	http.Handle("/custom", handle(bar, Notify(logger, "1"), Notify(logger, "2")))
	http.ListenAndServe(":80", nil)
}	