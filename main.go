package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(responseWriter http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(responseWriter, "Welcome to the HomePage!")
}

func main() {
	http.HandleFunc("/hello", hello)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
