// Third-party ServeMux - Julien Schmidt's Router
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	mux := httprouter.New()
	mux.GET("/", index)
	mux.GET("/about", about)
	mux.GET("/contact", contact)
	mux.GET("/apply", apply)
	mux.POST("/apply", applyProcess)
	mux.GET("/user/:name", user)
	mux.GET("/blog/:category/:article", blogRead)
	mux.POST("/blog/:category/:article", blogWrite)
	http.ListenAndServe(":8080", mux)
}

func index(rw http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(rw, "index.gohtml", nil)
	HandleError(rw, err)
}

func about(rw http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(rw, "about.gohtml", nil)
	HandleError(rw, err)
}

func contact(rw http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(rw, "contact.gohtml", nil)
	HandleError(rw, err)
}

func user(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	//err := tpl.ExecuteTemplate(rw, "user.gohtml", nil)
	//HandleError(rw, err)
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Julien Schemidt Router for Golang</title></head>
	<body>
	<h3>Hello %s! Thanks for using Julien Schmidt Router!!</h3><br>
	<a href="/">return</a><br>
	</body></html>`

	fmt.Fprintf(rw, body, ps.ByName("name"))

}

func apply(rw http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(rw, "apply.gohtml", nil)
	HandleError(rw, err)
}

func applyProcess(rw http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(rw, "applyProcess.gohtml", nil)
	HandleError(rw, err)
}

func blogRead(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Julien Schemidt Router for Golang</title></head>
	<body>
	<h3>READ CATEGORY: %s!</h3>
	<h3>READ ARTICLE: %s!</h3><br>
	<a href="/">return</a><br>
	</body></html>`

	fmt.Fprintf(rw, body, ps.ByName("category"), ps.ByName("article"))
	// fmt.Fprintf(rw, "READ CATEGORY, %s!\n", ps.ByName("category"))
	// fmt.Fprintf(rw, "READ ARTICLE, %s!\n", ps.ByName("article"))
	// fmt.Fprintln(rw, `<a href="/">return</a>`)
}

func blogWrite(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(rw, "WRITE CATEGORY, %s!\n", ps.ByName("category"))
	fmt.Fprintf(rw, "WRITE ARTICLE, %s!\n", ps.ByName("article"))
	fmt.Fprintln(rw, `<a href="/">return</a>`)
}

func update(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(rw, "UPDATE, %s!\n", ps.ByName("name"))
}

func trailing(rw http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(rw, "TRAILING, %s!\n", ps.ByName("name"))
}

func HandleError(rw http.ResponseWriter, err error) {
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	}
}
