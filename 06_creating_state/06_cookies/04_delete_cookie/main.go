package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.ListenAndServe(":8080", nil)
}

func index(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, `<h1><a href="/set">set a cookie</a></h1>`)
}

func set(rw http.ResponseWriter, req *http.Request) {
	http.SetCookie(rw, &http.Cookie{
		Name:  "session",
		Value: "some value",
	})
	fmt.Fprintln(rw, `<h1><a href="/read">read</a></h1>`)
}

func read(rw http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		http.Redirect(rw, req, "/set", http.StatusSeeOther)
		return
	}
	fmt.Fprintf(rw, `<h1>Your Cookie:<br>%v</h1><h1><a href="/expire">expire</a></h1>`, cookie)
}

func expire(rw http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		http.Redirect(rw, req, "/set", http.StatusSeeOther)
		return
	}
	cookie.MaxAge = -1 // delete cookie
	http.SetCookie(rw, cookie)
	http.Redirect(rw, req, "/", http.StatusSeeOther)
}
