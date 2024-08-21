package main

import (
	"fmt"
)

var print = fmt.Print
var println = fmt.Println
var printf = fmt.Printf

func getSum(nums ...int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

func main() {
	println(getSum(1, 2, 3, 4, 5, 6))
}
