package main

import (
	"fmt"
)

var print = fmt.Print
var println = fmt.Println
var printf = fmt.Printf

type Animate interface {
	AngrySound()
	HappySound()
}

type Cat string

func (cat Cat) Attack() {
	println("Cat attacks its prey")
}

func (cat Cat) Name() string {
	return string(cat)
}

func (cat Cat) AngrySound() {
	println("Cat says hissssssssss")
}

func (cat Cat) HappySound() {
	println("Cat says purrrrrrrrrr")
}

func main() {
	var kitty1 Animate

	kitty1 = Cat("Kitty1")
	kitty1.AngrySound()
	kitty1.HappySound()

	println()

	var kitty2 Cat
	kitty2 = kitty1.(Cat)
	kitty2.Attack()
	println(kitty2)
}
