package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
	// if there is an index.html located at the "/" directory
	// index.html will be automatically loaded instead of the
	// fileserver loading the directory
}
