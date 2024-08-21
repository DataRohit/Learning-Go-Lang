package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	fx := 1
	for fx <= 5 {
		pl(fx)
		fx++
	}
}
