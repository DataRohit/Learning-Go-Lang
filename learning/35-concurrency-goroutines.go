package main

import (
	"fmt"
	"time"
)

var print = fmt.Print
var println = fmt.Println
var printf = fmt.Printf

func printTo15() {
	for i := 1; i <= 15; i++ {
		println("Function 1 Value :", i)
	}
}

func printTo10() {
	for i := 1; i <= 10; i++ {
		println("Function 2 Value : ", i)
	}
}

func main() {
	go printTo15()
	go printTo10()

	time.Sleep(2 * time.Second)
}
