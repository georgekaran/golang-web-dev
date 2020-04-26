package main

import (
	"fmt"
	"net/http"
)

func handleRoute(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "You are in the index route!")
	case "/dog/":
		fmt.Fprintf(w, "You are in the dog route!")
	case "/me/":
		fmt.Fprintf(w, "My name is George!")
	default:
		fmt.Fprintf(w, "404 Page not found!")
	}
}

func main() {
	http.HandleFunc("/", handleRoute)
	http.HandleFunc("/dog/", handleRoute)
	http.HandleFunc("/me/", handleRoute)
	http.ListenAndServe(":8080", nil)
}
