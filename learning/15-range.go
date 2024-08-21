package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	nums := []int{1, 2, 3}

	for _, num := range nums {
		pl(num)
	}
}
