package app2

import (
	"fmt"
	"regexp"
)

var print = fmt.Print
var println = fmt.Println
var printf = fmt.Printf

func IsEmail(email string) (string, error) {
	regex, _ := regexp.Compile(`[\w._%+-]{1,20}@[\w.-]{2,20}.[A-Za-z]{2,3}`)

	if regex.MatchString(email) {
		return "Valid Email", nil
	} else {
		return "", fmt.Errorf("not a valid email")
	}
}
