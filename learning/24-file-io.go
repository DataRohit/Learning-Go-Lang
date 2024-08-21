package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var print = fmt.Print
var println = fmt.Println
var printf = fmt.Printf

func main() {
	file, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	primeArr := []int{2, 3, 5, 7, 11}
	var strPrimeArr []string
	for _, i := range primeArr {
		strPrimeArr = append(strPrimeArr, strconv.Itoa(i))
	}

	for _, num := range strPrimeArr {
		_, err := file.WriteString(num + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	file, err = os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		println(scan.Text())
	}

	if err := scan.Err(); err != nil {
		log.Fatal()
	}
}
