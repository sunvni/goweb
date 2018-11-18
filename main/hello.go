package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("src/github.com/sunvni/goweb/templates/*.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)
}
func index(w http.ResponseWriter, r *http.Request) {
	name := "Hoang"
	tpl.ExecuteTemplate(w, "index.html", name)
}
