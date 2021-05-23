package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	//sages := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad", "LaoTze", "Confusius"}
	sages := map[string]string{
		"India":      "Gandhi",
		"American":   "MLK",
		"Meditate":   "Buddha",
		"Love":       "Jesus",
		"Prophet":    "Muhammad",
		"Philosophy": "Lao Tze",
	}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
