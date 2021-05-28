package main

import (
	"io"
	"net/http"
)

type hotdog int

func (hd hotdog) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(rw, "doggy doggy doggy")
	case "/cat":
		io.WriteString(rw, "kitty kitty kitty")
	}
}

func main() {
	var hd hotdog
	http.ListenAndServe(":8080", hd)
}
