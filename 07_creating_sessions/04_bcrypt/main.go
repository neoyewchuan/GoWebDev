package main

import (
	"net/http"
	"text/template"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName  string
	Password  []byte
	FirstName string
	LastName  string
}

var tpl *template.Template
var dbUsers = map[string]user{}      // user ID, user
var dbSessions = map[string]string{} //session ID, user ID

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(rw http.ResponseWriter, req *http.Request) {
	u := getUser(rw, req)
	tpl.ExecuteTemplate(rw, "index.gohtml", u)
}

func bar(rw http.ResponseWriter, req *http.Request) {
	u := getUser(rw, req)
	if !alreadyLoggedIn(req) {
		http.Redirect(rw, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(rw, "bar.gohtml", u)
}

func signup(rw http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(rw, req, "/", http.StatusSeeOther)
		return
	}

	// process form submission
	if req.Method == http.MethodPost {

		// get form values
		uname := req.FormValue("username")
		passw := req.FormValue("password")
		fname := req.FormValue("firstname")
		lname := req.FormValue("lastname")

		// username taken?
		if _, ok := dbUsers[uname]; ok {
			http.Error(rw, "Username already taken", http.StatusForbidden)
			return
		}

		// create session
		sID := uuid.New()
		cookie := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(rw, cookie)
		dbSessions[cookie.Value] = uname

		// store user in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(passw), bcrypt.MinCost)
		if err != nil {
			http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		dbUsers[uname] = user{
			UserName:  uname,
			Password:  bs,
			FirstName: fname,
			LastName:  lname,
		}

		// redirect
		http.Redirect(rw, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(rw, "signup.gohtml", nil)
}
