package main

import (
	"fmt"
	"net/http"
)

func homefunc(rw http.ResponseWriter, rq *http.Request) {
	fmt.Fprintln(rw, "You are at the home page!")
}

func dogfunc(rw http.ResponseWriter, rq *http.Request) {
	fmt.Fprintln(rw, "dog doggy doggies everywhere!")
}

func mefunc(rw http.ResponseWriter, rq *http.Request) {
	fmt.Fprintln(rw, "Neo Yew Chuan")
}

func main() {

	http.HandleFunc("/", homefunc)
	http.HandleFunc("/dog/", dogfunc)
	http.HandleFunc("/me/", mefunc)

	http.ListenAndServe(":8080", nil)
}
