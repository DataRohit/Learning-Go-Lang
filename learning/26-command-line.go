package main

import (
	"fmt"
	"os"
	"strconv"
)

var print = fmt.Print
var println = fmt.Println
var printf = fmt.Printf

func main() {
	println(os.Args)

	args := os.Args[1:]

	var intArgs = []int{}
	for _, val := range args {
		val, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}

		intArgs = append(intArgs, val)
	}

	maxVal := 0
	for _, val := range intArgs {
		maxVal = max(maxVal, val)
	}
	println(maxVal)
}
