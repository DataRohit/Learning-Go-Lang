package main

import (
	"fmt"
)

var print = fmt.Print
var println = fmt.Println
var printf = fmt.Printf

func main() {
	var arr1 [5]int
	arr1[0] = 1

	for _, num := range arr1 {
		print(num, " ")
	}

	println("\n")

	arr2 := [5]int{1, 2, 3, 4, 5}
	println(arr2[0])
	println(len(arr2))

	print("\n")

	for i := 0; i < len(arr2); i++ {
		print(arr2[i], " ")
	}

	println("\n")

	arr3 := [2][2]int{
		{1, 2},
		{3, 4},
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			print(arr3[i][j], " ")
		}
		print("\n")
	}
}
