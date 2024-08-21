package stuff

import (
	"errors"
	"strconv"
	"time"
)

var Name string = "Rohit"

func IntArrToStrArr(intArr []int) []string {
	var strArr []string
	for _, val := range intArr {
		strArr = append(strArr, strconv.Itoa(val))
	}
	return strArr
}

type Date struct {
	day   int
	month int
	year  int
}

func (date *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("incorrect day value")
	}
	date.day = day
	return nil
}

func (date *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("incorrect month value")
	}
	date.month = month
	return nil
}

func (date *Date) SetYear(year int) error {
	if year < 1875 || year > time.Now().Year() {
		return errors.New("incorrect year value")
	}
	date.year = year
	return nil
}

func (date *Date) Day() int {
	return date.day
}

func (date *Date) Month() int {
	return date.month
}

func (date *Date) Year() int {
	return date.year
}
