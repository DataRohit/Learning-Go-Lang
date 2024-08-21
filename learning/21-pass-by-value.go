package main

import (
	"fmt"
)

var print = fmt.Print
var println = fmt.Println
var printf = fmt.Printf

func changeVal(num int) int {
	num++
	return num
}

func main() {
	num := 5
	println("Before Function:", num)
	changeVal(num)
	println("After Function:", num)
}
