package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(rw http.ResponseWriter, req *http.Request) {
	http.SetCookie(rw, &http.Cookie{
		Name:  "my-cookie",
		Value: "some dummy value",
	})
	fmt.Println(rw, "COOKIE WRITTEN : CHECK YOUR BROWSER")
	fmt.Println(rw, "i n chrome go to dev tools / application / cookies")
}

func read(rw http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Println(rw, "YOUR COOKIE:", c)
}
