package main

import (
	"html/template"
	"net/http"

	"github.com/google/uuid"
)

// for this code to run, you will need this package:
// go get github.com/satori/go.uuid

type user struct {
	UserName  string
	FirstName string
	LastName  string
}

var tpl *template.Template
var dbUsers = map[string]user{}      // user ID, user
var dbSessions = map[string]string{} // session ID, user ID

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(rw http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		id := uuid.New()
		// if err != nil {
		// 	fmt.Printf("Something went wrong: %s", err)
		// 	return
		// }
		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			//Secure: true,
			//HttpOnly: true,
		}
		http.SetCookie(rw, cookie)
	}

	// if the user exists already, get user
	var us user
	if un, ok := dbSessions[cookie.Value]; ok {
		us = dbUsers[un]
	}

	// process form submissions
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		fn := req.FormValue("firstname")
		ln := req.FormValue("lastname")
		us = user{
			UserName:  un,
			FirstName: fn,
			LastName:  ln, //
		}
		dbSessions[cookie.Value] = un
		dbUsers[un] = us
	}

	tpl.ExecuteTemplate(rw, "index.gohtml", us)
}

func bar(rw http.ResponseWriter, req *http.Request) {

	// get cookies
	cookie, err := req.Cookie("session")
	if err != nil {
		http.Redirect(rw, req, "/", http.StatusSeeOther)
		return
	}
	un, ok := dbSessions[cookie.Value]
	if !ok {
		http.Redirect(rw, req, "/", http.StatusSeeOther)
		return
	}
	us := dbUsers[un]
	tpl.ExecuteTemplate(rw, "bar.gohtml", us)
}

// https://play.golang.org/p/OKGL6phY_x
// http://play.golang.org/p/yORyGUZviV
