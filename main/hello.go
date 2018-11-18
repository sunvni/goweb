package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("src/github.com/sunvni/goweb/templates/*.html"))
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login/", loginHandler)
	http.HandleFunc("/home/", homeHandler)
	http.HandleFunc("/upload/", uploadHandler)
	http.Handle("/assets/css/", http.StripPrefix("/assets/css/", http.FileServer(http.Dir("src/github.com/sunvni/goweb/assets/css"))))
	http.Handle("/assets/js/", http.StripPrefix("/assets/js/", http.FileServer(http.Dir("src/github.com/sunvni/goweb/assets/js"))))
	http.ListenAndServe(":3000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))
	tpl.ExecuteTemplate(w, "login.html", token)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if len(r.FormValue("token")) == 0 {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
		r.ParseForm()
		fmt.Println(template.HTMLEscapeString(r.FormValue("username")))
		fmt.Println(r.Form.Get("username"))
		fmt.Println(template.HTMLEscapeString(r.FormValue("password")))
		http.Redirect(w, r, "/home/", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "home.html", nil)
}

func uploadHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		http.Redirect(w, req, "/home/", http.StatusSeeOther)
	} else {
		req.ParseMultipartForm(32 << 20)
		file, handler, err := req.FormFile("avatar")
		if err != nil {
			fmt.Print(err)
		}

		defer file.Close()

		fmt.Fprintf(w, "%v", handler.Header)

		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

		if err != nil {
			fmt.Print(err)
			return
		}

		defer f.Close()
		io.Copy(f, file)
	}
}
