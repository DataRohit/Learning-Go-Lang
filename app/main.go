package main

import (
	"fmt"
	"log"
	"reflect"

	stuff "example.com/project/mypackage"
)

func main() {
	fmt.Println("Hello", stuff.Name)
	intArr := []int{2, 3, 5, 7, 11}
	strArr := stuff.IntArrToStrArr(intArr)
	fmt.Println(strArr)
	fmt.Println(reflect.TypeOf(strArr))

	fmt.Println()

	date := stuff.Date{}
	if err := date.SetDay(21); err != nil {
		log.Fatal(err)
	}
	if err := date.SetMonth(12); err != nil {
		log.Fatal(err)
	}
	if err := date.SetYear(1974); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("1st Day : %d/%d/%d\n", date.Day(), date.Month(), date.Year())
}
