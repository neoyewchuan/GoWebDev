package main

import (
	"net/http"
)

func getUser(rw http.ResponseWriter, req *http.Request) user {
	var u user

	// get cookies
	cookie, err := req.Cookie("session")
	if err != nil {
		return u
	}

	// if the user exist already, get user

	if uname, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[uname]
	}
	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	cookie, err := req.Cookie("session")
	if err != nil {
		return false
	}
	uname := dbSessions[cookie.Value]
	_, ok := dbUsers[uname]
	return ok
}
