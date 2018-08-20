package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func handlerFunc1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my awesome second site!</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/test", handlerFunc1)
	http.ListenAndServe(":3000", nil)
}
