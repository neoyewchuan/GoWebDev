package main

import (
	"io"
	"net/http"
)

type hotdog int

func (hd hotdog) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	io.WriteString(rw, "dog doggy doggies")
}

type hotcat int

func (hc hotcat) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	io.WriteString(rw, "cat kitty kitties")
}

func main() {
	var hd hotdog
	var hc hotcat

	mux := http.NewServeMux()
	mux.Handle("/dog/", hd)
	mux.Handle("/cat", hc)

	http.ListenAndServe(":8080", mux)

	// can also use the default servemux
	// remove line 26 and change line 27 ~ 30 to the follow
	// http.Handle("/dog/", hd)
	// http.Handle("/cat", hc)
	//
	// http.ListenAndServe(":8080", nil)
}
