package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

type incrementer struct {
	Start     int `json:"start"`
	Increment int `json:"increment"`
}

func increment(w http.ResponseWriter, r *http.Request) {
	var i incrementer
	json.NewDecoder(r.Body).Decode(&i)
	_, _ = fmt.Fprintf(w, strconv.Itoa(i.Start+i.Increment))
}

func createRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello/{name}", greeting)
	r.HandleFunc("/hello", greeting)
	r.HandleFunc("/goodbye", goodbye)
	r.HandleFunc("/incrementer", increment)
	return r
}

func main() {
	http.ListenAndServe(":8080", createRouter())
}
