package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundance", abundance)
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

	c1, err := req.Cookie("my-cookie")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(rw, "YOUR COOKIE #1:", c1)
	}
	c2, err := req.Cookie("general")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(rw, "YOUR COOKIE #2:", c2)
	}
	c3, err := req.Cookie("specific")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(rw, "YOUR COOKIE #3:", c3)
	}
}

func abundance(rw http.ResponseWriter, req *http.Request) {
	http.SetCookie(rw, &http.Cookie{
		Name:  "general",
		Value: "some other value about general thing",
	})
	http.SetCookie(rw, &http.Cookie{
		Name:  "specific",
		Value: "some other value about specific thing",
	})
	fmt.Println(rw, "COOKIE WRITTEN : CHECK YOUR BROWSER")
	fmt.Println(rw, "i n chrome go to dev tools / application / cookies")
}
