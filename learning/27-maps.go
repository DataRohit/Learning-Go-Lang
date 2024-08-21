package main

import (
	"fmt"
)

var print = fmt.Print
var println = fmt.Println
var printf = fmt.Printf

func main() {
	heros := make(map[string]string)
	villians := make(map[string]string)

	heros["Batman"] = "Bruce Wayne"
	heros["Superman"] = "Clark Kent"
	heros["The Flash"] = "Barry Allen"

	villians["Lex Luther"] = "Alexander Joseph Luther"
	villians["Joker"] = "Jack Oswald White"

	superPets := map[int]string{
		1: "Krypto",
		2: "Bat Hound",
	}

	println(heros)
	println(villians)
	println(superPets)

	println()

	printf("Batman is %v\n", heros["Batman"])
	printf("Chip is %v\n", superPets[3])

	println()

	_, ok := superPets[3]
	println("Is there a 3rd pet :", ok)

	println()

	for key, val := range heros {
		printf("%s is %s\n", key, val)
	}

	println()

	delete(heros, "The Flash")
	println(heros)
}
