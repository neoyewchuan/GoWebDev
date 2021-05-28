package main

import (
	"io"
	"net/http"
)

func dogfunc(rw http.ResponseWriter, req *http.Request) {

	io.WriteString(rw, "dog doggy doggies")
}

func catfunc(rw http.ResponseWriter, req *http.Request) {

	io.WriteString(rw, "cat kitty kitties")
}

func main() {

	http.HandleFunc("/dog/", dogfunc)
	http.HandleFunc("/cat", catfunc)

	http.ListenAndServe(":8080", nil)
}
