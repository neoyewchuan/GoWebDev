package main

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

// for this code to run, you will need this package:
// go get github.com/satori/go.uuid

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(rw http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session-id")
	if err != nil {
		id := uuid.Must(uuid.NewV4(), nil)
		// if err != nil {
		// 	fmt.Printf("Something went wrong: %s", err)
		// 	return
		// }
		cookie = &http.Cookie{
			Name:  "session-id",
			Value: id.String(),
			//Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(rw, cookie)
	}
	fmt.Println(cookie)
}
