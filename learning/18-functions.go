package main

import (
	"fmt"
)

var print = fmt.Print
var println = fmt.Println
var printf = fmt.Printf

func getSum(x int, y int) int {
	return x + y
}

func getSumDiff(x int, y int) (int, int) {
	return x + y, x - y
}

func getQuotient(x float64, y float64) (ans float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("You can;'t divide by zero")
	} else {
		return x / y, nil
	}
}

func main() {
	println(getSum(5, 6))
	println(getSumDiff(4, 3))
	println(getQuotient(5, 0))
	println(getQuotient(22, 7))
}
