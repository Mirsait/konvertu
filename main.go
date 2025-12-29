package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := ViewData{
			Title: "length",
			Units: []string{
				"millimeter", "centimeter", "meter", "kilometer",
				"foot", "inch", "yard", "mile"},
		}
		tmpl, _ := template.ParseFiles("templates/index.html")
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/temperature", func(w http.ResponseWriter, r *http.Request) {
		data := ViewData{
			Title: "temperature",
			Units: []string{"celcius", "kelvin", "fahrenheit"},
		}
		tmpl, _ := template.ParseFiles("templates/index.html")
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/weight", func(w http.ResponseWriter, r *http.Request) {
		data := ViewData{
			Title: "weight",
			Units: []string{"milligram", "gram", "kilogram", "ounce", "pound"},
		}
		tmpl, _ := template.ParseFiles("templates/index.html")
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/result", func(w http.ResponseWriter, r *http.Request) {
		fromUnit := r.FormValue("fromUnit")
		toUnit := r.FormValue("toUnit")
		value := r.FormValue("value")
		prevPage := r.FormValue("page")
		data := ResultData{
			Title:     prevPage,
			FromUnit:  fromUnit,
			ToUnit:    toUnit,
			FromValue: value,
			ToValue:   "100",
		}
		tmpl, _ := template.ParseFiles("templates/result.html")
		tmpl.Execute(w, data)
	})

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:3000", nil)
}
