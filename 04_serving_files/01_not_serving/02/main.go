package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(rw http.ResponseWriter, req *http.Request) {

	rw.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(rw, `
	<!--image doesn't serve -->
	<img src="/toby.jpg">
	`)
}
