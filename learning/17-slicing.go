package main

import (
	"fmt"
)

var print = fmt.Print
var println = fmt.Println
var printf = fmt.Printf

func main() {
	slice1 := make([]string, 6)

	slice1[0] = "Society"
	slice1[1] = "of"
	slice1[2] = "the"
	slice1[3] = "Simulated"
	slice1[4] = "Universe"

	for i := 0; i < len(slice1); i++ {
		println(slice1[i])
	}

	for _, x := range slice1 {
		println(x)
	}

	sliceArr := [5]int{1, 2, 3, 4, 5}
	slice2 := sliceArr[0:2]
	println(slice2)

	slice2 = append(slice2, 7)
	println(slice2)
}
