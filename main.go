package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func greet(name string) string {
	if name == "" {
		return "hello"
	}
	return "hello " + name
}

func greeting(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, greet(mux.Vars(r)["name"]))
}

func goodbye(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, "goodbye")
}

func createRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello/{name}", greeting)
	r.HandleFunc("/hello", greeting)
	r.HandleFunc("/goodbye", goodbye)
	return r
}

func main() {
	http.ListenAndServe(":8080", createRouter())
}
