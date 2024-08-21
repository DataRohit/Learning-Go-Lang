package main

import (
	"fmt"
)

var print = fmt.Print
var println = fmt.Println
var printf = fmt.Printf

type rectangle struct {
	length, height float64
}

func (rect rectangle) Area() float64 {
	return rect.length * rect.height
}

func main() {
	rect1 := rectangle{
		length: 10,
		height: 15,
	}
	println(rect1.Area())
}
