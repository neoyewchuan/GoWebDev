package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(rw, `
	<img src="toby.jpg">
	`)
}

func dogPic(rw http.ResponseWriter, req *http.Request) {
	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(rw, "file not found", http.StatusNotFound)
	}
	defer f.Close()

	io.Copy(rw, f)
}
