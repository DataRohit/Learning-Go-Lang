package main

import (
	"fmt"
)

var print = fmt.Print
var println = fmt.Println
var printf = fmt.Printf

func getArraySum(arr []int) int {
	sum := 0

	for _, num := range arr {
		sum += num
	}

	return sum
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	println(getArraySum(arr))
}
