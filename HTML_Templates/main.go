package main

import (
	"html/template"
	"net/http"
)

//var tpl *template.Template

var tpl *template.Template

func main() {

	tpl, _ = template.ParseFiles("index.html")

	// Handle requests for static files (e.g., CSS)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("."))))

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// func (t *Template) Execute(wr io.Writer, data interface{}) error
	tpl.Execute(w, nil)
}
