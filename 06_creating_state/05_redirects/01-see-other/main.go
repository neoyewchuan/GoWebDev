package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(rw http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at foo: ", req.Method, "\n\n")
}

func bar(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at bar: ", req.Method)
	rw.Header().Set("Location", "/")
	rw.WriteHeader(http.StatusSeeOther)
}

func barred(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at barred: ", req.Method)
	tpl.ExecuteTemplate(rw, "index.gohtml", nil)

}
