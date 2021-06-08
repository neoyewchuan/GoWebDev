package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/toby/", dog)
	//http.HandleFunc("/toby.jpg", toby)
	http.ListenAndServe(":8080", nil)

}

func foo(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "foo ran\n")
}

func dog(rw http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	tpl.ExecuteTemplate(rw, "dog.gohtml", nil)
}

// func toby(rw http.ResponseWriter, req *http.Request) {
// 	http.ServeFile(rw, req, "toby.jpg")
// }
