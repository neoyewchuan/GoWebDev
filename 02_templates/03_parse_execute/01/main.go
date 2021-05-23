package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("./../template/tpl.gohtml")
	if err != nil {
		log.Fatal("error parsing tpl.gohtml: ", err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal("error executing tpl.gohtml: ", err)
	}
}
