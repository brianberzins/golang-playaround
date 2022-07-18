package main

import (
	"fmt"
	"log"
	"net/http"
)

func greeting(responseWriter http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(responseWriter, "hello")
}

func goodbye(responseWriter http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(responseWriter, "goodbye")
}

func main() {
	http.HandleFunc("/hello", greeting)
	http.HandleFunc("/goodbye", goodbye)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
