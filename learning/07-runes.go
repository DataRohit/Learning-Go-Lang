package main

import (
	"fmt"
	"unicode/utf8"
)

var pl = fmt.Println

func main() {
	rStr1 := "abcdefg"
	pl(utf8.RuneCountInString(rStr1))

	for i, runeVal := range rStr1 {
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	}
}
