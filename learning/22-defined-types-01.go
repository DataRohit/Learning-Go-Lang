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

func (tsp TeaSpn) ToMLs() MilLit {
	return MilLit(tsp * 4.92)
}

func (tbsp TblSpn) ToMLs() MilLit {
	return MilLit(tbsp * 14.79)
}

func main() {
	tsp1 := TeaSpn(3)
	printf("%.2f tsp = %.2f ml\n", tsp1, tsp1.ToMLs())

	tbsp1 := TblSpn(3)
	printf("%.2f tbsp = %.2f ml\n", tbsp1, tbsp1.ToMLs())
}
