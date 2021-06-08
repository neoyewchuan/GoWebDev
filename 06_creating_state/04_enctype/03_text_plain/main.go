package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*"))
}

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(rw http.ResponseWriter, req *http.Request) {

	// body
	bs := make([]byte, req.ContentLength)
	req.Body.Read(bs)
	body := string(bs)

	err := tpl.ExecuteTemplate(rw, "index.gohtml", body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
