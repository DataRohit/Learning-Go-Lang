package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var pl = fmt.Println

func main() {
	seedSecs := time.Now().Unix()

	random_generator := rand.New(rand.NewSource(seedSecs))

	randNum := random_generator.Intn(50) + 1

	fmt.Print("Guess a number between 1 and 50\n")
	pl("Random Number is :", randNum)

	for true {
		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		guess = strings.TrimSpace(guess)
		intGuess, err := strconv.Atoi(guess)

		if err != nil {
			log.Fatal(err)
		}

		if intGuess > randNum {
			pl("Pick a Lower Value")
		} else if intGuess < randNum {
			pl("Pick a Higher Value")
		} else {
			pl("You Guessed It!")
			break
		}
	}
}
