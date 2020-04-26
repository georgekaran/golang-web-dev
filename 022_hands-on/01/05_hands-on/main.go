package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type handlerMuxIndex int

func (h handlerMuxIndex) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", "Index yo!")
}

type handlerMuxDog int

func (h handlerMuxDog) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	tpl.ExecuteTemplate(w, "dog.gohtml", "Dog is crazy.")
}

type handlerMuxMe int

func (h handlerMuxMe) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	tpl.ExecuteTemplate(w, "me.gohtml", "George!")
}

func main() {
	var handlerDog handlerMuxDog
	var handlerMuxMe handlerMuxMe
	var handlerIndex handlerMuxIndex

	http.Handle("/", handlerIndex)
	http.Handle("/dog/", handlerDog)
	http.Handle("/me", handlerMuxMe)

	http.ListenAndServe(":8080", nil)
}
