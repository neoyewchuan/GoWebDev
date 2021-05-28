package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (hd hotdog) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Mcleod-Key", "this is from mcleod")
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	rw.WriteHeader(http.StatusOK)
	fmt.Fprintln(rw, "<h1>Any code you want in this func</h1>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
