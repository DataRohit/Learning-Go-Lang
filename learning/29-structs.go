package main

import (
	"fmt"
)

var print = fmt.Print
var println = fmt.Println
var printf = fmt.Printf

type customer struct {
	name    string
	address string
	balance float64
}

func getCustomerInfo(cust customer) {
	printf("%s owes us %.2f and they live at %s\n", cust.name, cust.balance, cust.address)
}

func updateCustomerAddress(cust *customer, address string) {
	cust.address = address
}

func main() {
	var cust1 customer
	cust1.name = "Tom Smith"
	cust1.address = "5 Main St"
	cust1.balance = 234.56

	getCustomerInfo(cust1)
	updateCustomerAddress(&cust1, "South Street")
	getCustomerInfo(cust1)

	cust2 := customer{
		"Sally Smith",
		"123 Main",
		0.0,
	}
	getCustomerInfo(cust2)
}
