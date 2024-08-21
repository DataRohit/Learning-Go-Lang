package main

import (
	"fmt"
)

var print = fmt.Print
var println = fmt.Println
var printf = fmt.Printf

type contact struct {
	fName string
	lName string
	phone string
}

type business struct {
	name    string
	address string
	contact
}

func (b business) info() {
	printf("Contact at %s is %s %s\n", b.name, b.contact.fName, b.contact.lName)
}

func main() {
	cont1 := contact{
		fName: "James",
		lName: "Wang",
		phone: "555-1212",
	}

	bus1 := business{
		name:    "ABC Plumbing",
		address: "234 North St",
		contact: cont1,
	}

	bus1.info()
}
