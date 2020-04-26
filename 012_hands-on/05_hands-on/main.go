package main

import (
	"html/template"
	"log"
	"os"
)

type food struct {
	Name string
	Price float64
}

type menu struct {
	Breakfast []food
	Lunch []food
	Dinner []food
}

var tpl *template.Template

func init() {
	tpl = template.Must(tpl.ParseFiles("tpl.gohtml"))
}

func main() {
	restaurant := menu{
		Breakfast: []food{
			{"Eggs", 2.0},
			{"Bacon", 4.0},
			{"Bread", 1.0},
		},
		Lunch: []food{
			{"Rice", 3.0},
			{"Beans", 2.0},
			{"Meat", 8.0},
		},
		Dinner: []food{
			{"Soup", 4.0},
			{"Wine", 10.0},
			{"Something", 20.0},
		},
	}

	err := tpl.Execute(os.Stdout, restaurant)
	if err != nil {
		log.Fatalln(err)
	}
}