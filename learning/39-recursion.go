package main

import (
	"fmt"
)

var print = fmt.Print
var println = fmt.Println
var printf = fmt.Printf

func factorial(num int) int {
	if num == 0 {
		return 1
	}

	return num * factorial(num-1)
}

func main() {
	var number int
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	println(factorial(number))
}
