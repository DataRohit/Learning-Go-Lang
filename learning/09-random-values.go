package main

import (
	"fmt"
	"math/rand"
	"time"
)

var pl = fmt.Println

func main() {
	seedSecs := time.Now().Unix()

	random_generator := rand.New(rand.NewSource(seedSecs))

	randNum := random_generator.Intn(50) + 1

	pl(randNum)
}
