package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type course struct {
	Number string
	Name   string
	Units  float64
}

type semester struct {
	Term    string
	Courses []course
}

type academicyear struct {
	Fall, Spring, Summer, Winter semester
}

func main() {

	fa1 := course{
		Number: "CSCI-40",
		Name:   "Introduction to Programming in Go",
		Units:  0.5,
	}
	fa2 := course{
		Number: "CSCI-130",
		Name:   "Introduction to Web Programming with Go",
		Units:  0.5,
	}
	fa3 := course{
		Number: "CSCI-140",
		Name:   "Mobile Apps Using Go",
		Units:  1,
	}
	fa4 := course{
		Number: "CSCI-200",
		Name:   "Introduction to Blockchain & Cryptocurrency",
		Units:  0.5,
	}
	sp1 := course{
		Number: "CSCI-50",
		Name:   "Advanced Go",
		Units:  1,
	}
	sp2 := course{
		Number: "CSCI-190",
		Name:   "Advanced Web Programming with Go",
		Units:  1,
	}
	sp3 := course{
		Number: "CSCI-191",
		Name:   "Advanced Mobile Apps with Go",
		Units:  1,
	}
	su1 := course{
		Number: "CSCI-202",
		Name:   "Blockchain Use Case",
		Units:  0.5,
	}

	su2 := course{
		Number: "CSCI-211",
		Name:   "Enterprise Blockchain Fundamental",
		Units:  0.5,
	}

	su3 := course{
		Number: "CSCI-230",
		Name:   "Ethereum Smart Contract with Solidity",
		Units:  1.0,
	}

	fs1 := semester{
		Term:    "Fall 2021",
		Courses: []course{fa1, fa2, fa3, fa4},
	}

	ss1 := semester{
		Term:    "Spring 2021",
		Courses: []course{sp1, sp2, sp3},
	}

	ss2 := semester{
		Term:    "Summer 2021",
		Courses: []course{su1, su2, su3},
	}

	ay := academicyear{
		Fall:   fs1,
		Spring: ss1,
		Summer: ss2,
	}

	err := tpl.Execute(os.Stdout, ay)
	if err != nil {
		log.Fatalln(err)
	}

}
