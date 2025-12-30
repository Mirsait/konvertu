package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := ViewData{
			Title: "length",
			Units: getNames(Length),
		}
		tmpl, _ := template.ParseFiles("templates/index.html")
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/temperature", func(w http.ResponseWriter, r *http.Request) {
		data := ViewData{
			Title: "temperature",
			Units: getNames(Temperature),
		}
		tmpl, _ := template.ParseFiles("templates/index.html")
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/weight", func(w http.ResponseWriter, r *http.Request) {
		data := ViewData{
			Title: "weight",
			Units: getNames(Mass),
		}
		tmpl, _ := template.ParseFiles("templates/index.html")
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/result", func(w http.ResponseWriter, r *http.Request) {
		fromUnit := r.FormValue("fromUnit")
		toUnit := r.FormValue("toUnit")
		value := r.FormValue("value")
		prevPage := r.FormValue("page")

		val, _ := strconv.ParseFloat(value, 64)

		result, _ := Convert(val, fromUnit, toUnit)

		data := ResultData{
			Title:     prevPage,
			FromUnit:  getSymbol(fromUnit),
			ToUnit:    getSymbol(toUnit),
			FromValue: fmt.Sprintf("%.2f", val),
			ToValue:   fmt.Sprintf("%.2f", result),
		}
		tmpl, _ := template.ParseFiles("templates/result.html")
		tmpl.Execute(w, data)
	})

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:3000", nil)
}
