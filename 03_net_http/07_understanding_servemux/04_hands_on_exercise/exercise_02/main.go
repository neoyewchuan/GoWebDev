package main

import (
	"net/http"
	"text/template"
)

func homefunc(rw http.ResponseWriter, rq *http.Request) {

	// err := rq.ParseForm()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	data := struct {
		Message string
	}{
		"This is the home page",
	}
	tpl.ExecuteTemplate(rw, "index.gohtml", data)
}

func dogfunc(rw http.ResponseWriter, rq *http.Request) {
	// err := rq.ParseForm()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// rw.Header().Set("MyKey", "dog doggy doggies everywhere!")
	// tpl.ExecuteTemplate(rw, "index.gohtml", rq.Form)
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	data := struct {
		Message string
	}{
		"dog doggy doggies everywhere!",
	}
	tpl.ExecuteTemplate(rw, "index.gohtml", data)

}

func mefunc(rw http.ResponseWriter, rq *http.Request) {
	// err := rq.ParseForm()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// rw.Header().Set("MyKey", "Neo Yew Chuan")
	// tpl.ExecuteTemplate(rw, "index.gohtml", rq.Form)
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	data := struct {
		Message string
	}{
		"Neo Yew Chuan",
	}
	tpl.ExecuteTemplate(rw, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {

	http.HandleFunc("/", homefunc)
	http.HandleFunc("/dog/", dogfunc)
	http.HandleFunc("/me/", mefunc)

	http.ListenAndServe(":8080", nil)
}
