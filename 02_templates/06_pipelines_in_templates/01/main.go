package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

func dayMonthYear(t time.Time) string {
	return t.Format("02-01-2006")
}

func yearMonthDay(t time.Time) string {
	return t.Format("2006-01-02")
}

func dayMonthYearTime(t time.Time) string {
	return t.Format("02-01-2006 03:04")
}

var fm = template.FuncMap{
	"fdateMDY":  monthDayYear,
	"fdateDMY":  dayMonthYear,
	"fdateYMD":  yearMonthDay,
	"fdateTime": dayMonthYearTime,
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
