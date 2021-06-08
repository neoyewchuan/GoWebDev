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
	bs, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	dbUsers["test@test.com"] = user{"test@test.com", bs, "James", "Bond"}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
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

func login(rw http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(rw, req, "/", http.StatusSeeOther)
		return
	}
	var u user
	// process form submissions
	if req.Method == http.MethodPost {
		uname := req.FormValue("username")
		passw := req.FormValue("password")
		// is there a username?
		u, ok := dbUsers[uname]
		if !ok {
			http.Error(rw, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// does the entered password match the stored password
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(passw))
		if err != nil {
			http.Error(rw, "Username and/or password do not match", http.StatusForbidden)
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
		http.Redirect(rw, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(rw, "login.gohtml", u)

}

func logout(rw http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(req) {
		http.Redirect(rw, req, "/login", http.StatusSeeOther)
		return
	}

	cookie, _ := req.Cookie("session")
	// delete the cookie
	delete(dbSessions, cookie.Value)
	//remove the cookies
	cookie = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(rw, cookie)
	http.Redirect(rw, req, "/login", http.StatusSeeOther)

}
