package main

import (
	"fmt"
)

var print = fmt.Print
var println = fmt.Println
var printf = fmt.Printf

func nums1(channel chan int) {
	channel <- 1
	channel <- 2
	channel <- 3
}

func nums2(channel chan int) {
	channel <- 4
	channel <- 5
	channel <- 6
}

func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)

	go nums1(channel1)
	go nums2(channel2)

	println(<-channel1)
	println(<-channel1)
	println(<-channel1)

	println(<-channel2)
	println(<-channel2)
	println(<-channel2)
}
