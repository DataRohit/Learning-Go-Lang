package main

import (
	"fmt"
	"regexp"
)

var print = fmt.Print
var println = fmt.Println
var printf = fmt.Printf

func main() {
	reStr1 := "The ape was at the apex"

	match1, _ := regexp.MatchString(`ape[^ ]`, reStr1)
	println("Match 1:", match1)

	match2, _ := regexp.MatchString(`ape\s`, reStr1)
	println("Match 2:", match2)

	match3, _ := regexp.MatchString(`apez`, reStr1)
	println("Match 3:", match3)

	match4, _ := regexp.MatchString(`^apex`, reStr1)
	println("Match 4:", match4)

	println()

	reStr2 := "Cat rat mat fat pat"
	r, _ := regexp.Compile("([crmfp]at)")
	println(r.MatchString(reStr2))
	println(r.FindString(reStr2))
	println(r.FindStringIndex(reStr2))
	println(r.FindAllString(reStr2, -1))
	println(r.FindAllString(reStr2, 2))
	println(r.FindAllStringSubmatchIndex(reStr2, -1))
	println(r.ReplaceAllString(reStr2, "Dog"))
}
