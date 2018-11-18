package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

//Product need a comment
type Product struct {
	Name        string
	Price       int
	Description string
}

func init() {
	tpl = template.Must(template.ParseGlob("src/github.com/sunvni/goweb/templates/*.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/assets/css/", http.StripPrefix("/assets/css/", http.FileServer(http.Dir("src/github.com/sunvni/goweb/assets/css"))))
	http.Handle("/assets/js/", http.StripPrefix("/assets/js/", http.FileServer(http.Dir("src/github.com/sunvni/goweb/assets/js"))))
	http.ListenAndServe(":3000", nil)
}

func assets(w http.ResponseWriter, r *http.Request) {

}

func index(w http.ResponseWriter, r *http.Request) {

	vest := Product{
		Name:        "Ao Vest",
		Price:       10000,
		Description: "Lam tu chat lieu vai bong",
	}
	aokhoac := Product{
		Name:        "Ao Khoac",
		Price:       7999,
		Description: "Lam tu chat lieu vai bong",
	}
	quan := Product{
		Name:        "Quan",
		Price:       2980,
		Description: "Lam tu chat lieu vai bong",
	}

	products := [3]Product{vest, aokhoac, quan}

	tpl.ExecuteTemplate(w, "index.html", products)
}
