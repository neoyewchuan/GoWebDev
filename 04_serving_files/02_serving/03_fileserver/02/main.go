package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func dog(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(rw, `<img src="/resources/toby.jpg">`)
	io.WriteString(rw, `<img src="/resources/henry.jpg">`)
	io.WriteString(rw, `<img src="/resources/beagle.jpg">`)

}
