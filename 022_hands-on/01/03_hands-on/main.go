package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func handleRoute(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		tpl.ExecuteTemplate(w, "index.gohtml", "Index yo!")
	case "/dog/":
		tpl.ExecuteTemplate(w, "dog.gohtml", "Dog is crazy.")
	case "/me/":
		tpl.ExecuteTemplate(w, "me.gohtml", "George!")
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
