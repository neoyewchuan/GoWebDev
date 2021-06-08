package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", set)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(rw http.ResponseWriter, req *http.Request) {
	counter := 0
	c, err := req.Cookie("visit-count")
	if err == http.ErrNoCookie {
		counter = 0
	} else {
		counter, err = strconv.Atoi(c.Value)
		if err != nil {
			counter = 0
		}
	}
	counter++
	c.Value = strconv.Itoa(counter)
	http.SetCookie(rw, c)
	fmt.Fprintln(rw, "You've visited this site "+strconv.Itoa(counter)+" times")
}
