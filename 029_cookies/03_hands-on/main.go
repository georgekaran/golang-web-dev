package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	cookieVisits, err := r.Cookie("visits")
	if err != nil {
		cookieVisits = &http.Cookie{
			Name: "visits",
			Value: "0",
			Path: "/",
		}
	}
	updateCookieValue(w, cookieVisits)

	fmt.Fprintf(w, "You've been %v time(s) here.", cookieVisits.Value)
}

func updateCookieValue(w http.ResponseWriter, c *http.Cookie) {
	currentValue, err := strconv.Atoi(c.Value)
	if err != nil {
		currentValue = 0
	}
	currentValue++
	c.Value = strconv.Itoa(currentValue)
	http.SetCookie(w, c)
}
