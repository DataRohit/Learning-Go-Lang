package main

import (
	"fmt"
)

var print = fmt.Print
var println = fmt.Println
var printf = fmt.Printf

func changeVal(num *int) {
	*num += 15
}

func main() {
	num := 5
	println("Before Function:", num)
	changeVal(&num)
	println("After Function:", num)

	println()

	var numPtr *int = &num
	println("Address to Number:", numPtr)
	println("Value at Address:", *numPtr)

	*numPtr = 35
	println("Updated Value at Address:", *numPtr)
}
