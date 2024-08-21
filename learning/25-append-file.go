package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var print = fmt.Print
var println = fmt.Println
var printf = fmt.Printf

func main() {
	_, err := os.Stat("data.txt")
	if errors.Is(err, os.ErrNotExist) {
		println("File Doesn't Exist")
	} else {
		file, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		if _, err := file.WriteString("13\n"); err != nil {
			log.Fatal(err)
		}
	}
}
