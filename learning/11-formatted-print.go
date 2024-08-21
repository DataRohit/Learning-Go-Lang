package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	var strVal string = "Stuff"
	var intVal int = 123
	var charVal rune = 'A'
	var floatVal float32 = 3.14
	var boolVal bool = true
	var base8Val = 1
	var base16Val = 1

	fmt.Printf("%s %d %c %f %t %o %x\n", strVal, intVal, charVal, floatVal, boolVal, base8Val, base16Val)

	fmt.Printf("%9f\n", 3.14)
	fmt.Printf("%.2f\n", 3.141592)
	fmt.Printf("%9.f\n", 3.141592)

	sp1 := fmt.Sprintf("%9.f\n", 3.141592)
	pl(sp1)
}
