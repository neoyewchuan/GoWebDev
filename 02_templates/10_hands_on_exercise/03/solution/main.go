/* 1. Create a data structure to pass to a template which
contains information about restaurant's menu including Breakfast, Lunch, and Dinner items */
package main

import (
	"log"
	"os"
	"text/template"
)

type item struct {
	Name  string // name of menu item
	Desc  string
	Price float64
}

type menu struct {
	Breakfast []item
	Lunch     []item
	Dinner    []item
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	b1 := item{
		Name:  "Ham and Egg",
		Desc:  "Chicken ham with a sunny shine up",
		Price: 20.0,
	}
	b2 := item{
		Name:  "Bacon and Cheese",
		Desc:  "Bacon top with mozzarella cheese",
		Price: 10.0,
	}
	b3 := item{
		Name:  "Tuna Salad",
		Desc:  "Vegetable salad with delicious tuna",
		Price: 15.0,
	}
	b4 := item{
		Name:  "Beef Soup",
		Desc:  "The best prime quality beef soup in town",
		Price: 25.0,
	}

	l1 := item{
		Name:  "Spicy Beef Barbecue",
		Desc:  "Perfectly barbecued spicy prime grade beef",
		Price: 20.0,
	}
	l2 := item{
		Name:  "Pork Barbecue",
		Desc:  "Perfectly barbecued pork with secret spices",
		Price: 10.0,
	}
	l3 := item{
		Name:  "Oven Chicken Barbecue",
		Desc:  "Perfectly barbecued chicken with secret spices",
		Price: 15.0,
	}
	l4 := item{
		Name:  "Pulled Beef Barbecue Burger",
		Desc:  "The best pulled beef burger in town",
		Price: 25.0,
	}

	d1 := item{
		Name:  "Spicy Pork Barbecue",
		Desc:  "Perfectly barbecued spicy pork with secret spices",
		Price: 20.0,
	}
	d2 := item{
		Name:  "Vegetable Pork Barbecue",
		Desc:  "Perfectly barbecued pork with capsicum/lettuce",
		Price: 10.0,
	}
	d3 := item{
		Name:  "Oven Pork Barbecue",
		Desc:  "Perfectly barbecued pork with special spices",
		Price: 15.0,
	}
	d4 := item{
		Name:  "Pulled Pork Barbecue",
		Desc:  "Barbecued pull pork with special spices",
		Price: 25.0,
	}

	theMenu := menu{
		Breakfast: []item{b1, b2, b3, b4},
		Lunch:     []item{l1, l2, l3, l4},
		Dinner:    []item{d1, d2, d3, d4},
	}

	err := tpl.Execute(os.Stdout, theMenu)
	if err != nil {
		log.Fatalln(err)
	}
}
