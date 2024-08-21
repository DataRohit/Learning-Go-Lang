package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var pl = fmt.Println

func main() {
	reader := bufio.NewReader(os.Stdin)

	pl("Please enter your age:")
	input, _ := reader.ReadString('\n')

	age, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		pl("Invalid input. Please enter a valid age.")
		return
	}

	if age >= 1 && age <= 18 {
		pl("Important Birthday")
	} else if age == 21 || age == 50 || age >= 65 {
		pl("Important Birthday")
	} else {
		pl("Not an Important Birthday")
	}
}
