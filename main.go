package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Got request / \n")
	tempFile := "main.html"
	tmpl, err := template.New(tempFile).ParseFiles(tempFile)
	if err != nil {
		panic(err)
	}

	tmpl.Execute(w, "Juan")
}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Got request /user \n")
	name := r.URL.Query().Get("name")
	tempFile := "main.html"
	tmpl, err := template.New(tempFile).ParseFiles(tempFile)
	if err != nil {
		panic(err)
	}

	tmpl.Execute(w, name)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/user", user)

	http.ListenAndServe(":5050", mux)
}
