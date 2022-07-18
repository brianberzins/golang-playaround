package main

import (
	"fmt"
	"log"
	"net/http"
)

func greeting() string {
	return "Hello"
}

func hello(responseWriter http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(responseWriter, greeting())
}

func main() {
	http.HandleFunc("/hello", hello)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
