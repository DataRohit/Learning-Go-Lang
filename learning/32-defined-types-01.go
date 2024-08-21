package main

import (
	"fmt"
)

var print = fmt.Print
var println = fmt.Println
var printf = fmt.Printf

type TeaSpn float64
type TblSpn float64
type MilLit float64

func TeaSpnToMilLit(tsp TeaSpn) MilLit {
	return MilLit(tsp * 4.92)
}

func TblSpnToMilLit(tsp TblSpn) MilLit {
	return MilLit(tsp * 14.79)
}

func main() {
	ml1 := MilLit(TeaSpn(3) * 4.92)
	printf("3 tsp = %.2f mL\n", ml1)

	ml2 := MilLit(TblSpn(3) * 14.79)
	printf("3 tbsp = %.2f mL\n", ml2)

	println()
	println(TeaSpn(2) + TeaSpn(4))
	println(TeaSpn(2) > TeaSpn(4))

	println()
	printf("3 tsp = %.2f mL\n", TeaSpnToMilLit(3))
	printf("3 tbsp = %.2f mL\n", TblSpnToMilLit(3))
}
