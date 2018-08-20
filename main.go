package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please send an email "+
		"to <a href=\"mailto:support@test.test\">"+
		"support@test.test</a>.")
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "This is a nice faq page!")
}

func sorry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>I'm afraid i can't do that for you dave... ):")
}

func main() {
	var nf http.Handler = http.HandlerFunc(sorry)
	r := mux.NewRouter()
	r.NotFoundHandler = nf
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	http.ListenAndServe(":3000", r)
}
