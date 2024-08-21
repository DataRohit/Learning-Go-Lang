package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var print = fmt.Print
var println = fmt.Println
var printf = fmt.Printf

type MyConstraint interface {
	int | float64
}

func getSumGen[T MyConstraint](num1 T, num2 T) T {
	return num1 + num2
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter first number: ")
	input1, _ := reader.ReadString('\n')
	input1 = input1[:len(input1)-1]

	fmt.Print("Enter second number: ")
	input2, _ := reader.ReadString('\n')
	input2 = input2[:len(input2)-1]

	if num1, err := strconv.Atoi(input1); err == nil {
		if num2, err := strconv.Atoi(input2); err == nil {
			println("Sum (int):", getSumGen(num1, num2))
			return
		}
	}

	if num1, err := strconv.ParseFloat(input1, 64); err == nil {
		if num2, err := strconv.ParseFloat(input2, 64); err == nil {
			println("Sum (float64):", getSumGen(num1, num2))
			return
		}
	}

	println("Invalid input. Please enter valid numbers.")
}
