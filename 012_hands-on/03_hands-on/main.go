package main

import (
	"html/template"
	"log"
	"os"
)

type region struct {
	Name string
}

type hotels struct {
	Name string
	Address string
	City string
	Zip string
	Region region
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	regions := [3]region{
		{"Southern"},
		{"Central"},
		{"Northern"},
	}

	hotels := []hotels{
		{"Hotel 1", "Address 1", "New York", "92101-102", regions[0]},
		{"Hotel 2", "Address 2", "Los Angeles", "85122-102", regions[1]},
		{"Hotel 3", "Address 3", "Chicago", "421515-102", regions[2]},
	}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
