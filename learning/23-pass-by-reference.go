package main

import (
	"fmt"
)

var print = fmt.Print
var println = fmt.Println
var printf = fmt.Printf

func doubleValues(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] *= 2
	}
}

func getAverage(nums ...float64) float64 {
	var sum float64 = 0.0
	var numSize float64 = float64(len(nums))

	for _, val := range nums {
		sum += val
	}

	return sum / numSize
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	doubleValues(&nums)
	println(nums)

	slice := []float64{11, 13, 17}
	printf("%.4f", getAverage(slice...))
	println()
}
