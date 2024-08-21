package main

import (
	"fmt"
)

var print = fmt.Print
var println = fmt.Println
var printf = fmt.Printf

func useFunc(fn func(int, int) int, num1 int, num2 int) {
	println((fn(num1, num2)))
}

func sumVals(num1, num2 int) int {
	return num1 + num2
}

func main() {
	intSum := func(num1 int, num2 int) int { return num1 + num2 }
	println(intSum(5, 6))

	println()
	sample1 := 1
	println(sample1)
	changeVar := func() {
		sample1++
	}
	changeVar()
	println(sample1)

	println()
	useFunc(sumVals, 5, 8)
}
