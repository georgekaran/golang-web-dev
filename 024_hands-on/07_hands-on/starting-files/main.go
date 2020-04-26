package main

import (
	"log"
	"net/http"
	"html/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(tpl.ParseFiles("templates/index.gohtml"))
}

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/resources/", http.StripPrefix("/resources", fs))
	http.HandleFunc("/", dogs)
	http.ListenAndServe(":8080", nil)
}

func dogs(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}
