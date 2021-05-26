/* 1. Create a data structure to pass to a template which
contains information about California hotels including Name, Address, City, Zip, Region
region can be: Southern, Central, Northern
can hold an unlimited number of hotels
*/

package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	h1 := hotel{
		Name:    "Hampton Inn Oakhurst-Yosemite",
		Address: "40740 Highway 41",
		City:    "Oakhurst",
		Zip:     "93644",
		Region:  "Southern",
	}
	h2 := hotel{
		Name:    "Hampton Inn & Suites Fresno",
		Address: "327 E. Fir Avenue",
		City:    "Fresno",
		Zip:     "93720",
		Region:  "Central",
	}
	h3 := hotel{
		Name:    "Homewood Suites by Hilton Fresno",
		Address: "6820 North Fresno Street",
		City:    "Fresno",
		Zip:     "93721",
		Region:  "Central",
	}
	h4 := hotel{
		Name:    "Hilton Garden Inn Clovis",
		Address: "520 West Shaw Ave",
		City:    "Clovis",
		Zip:     "93612",
		Region:  "Southern",
	}
	h5 := hotel{
		Name:    "Hampton Inn & Suites Merced",
		Address: "225 South Parsons Avenue",
		City:    "Merced",
		Zip:     "95340",
		Region:  "Northern",
	}
	h6 := hotel{
		Name:    "Home2 Suites by Hilton Hanford Lemoore",
		Address: "1589 Glendale Avenue",
		City:    "Hanford",
		Zip:     "93230",
		Region:  "Southern",
	}
	h7 := hotel{
		Name:    "Hampton Inn & Suites Tulare",
		Address: "1100 N. Cherry Street",
		City:    "Tulare",
		Zip:     "93274",
		Region:  "Central",
	}

	hotels := []hotel{h1, h2, h3, h4, h5, h6, h7}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Println(err)
	}

}
