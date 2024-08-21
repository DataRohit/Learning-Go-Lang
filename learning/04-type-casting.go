package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var pl = fmt.Println

func main() {
	var1 := 3.5
	var2 := int(var1)
	pl(var2, reflect.TypeOf(var2))

	var3 := "500000000000000"
	var4, err := strconv.Atoi(var3)
	pl(var4, err, reflect.TypeOf(var4))

	var5 := 500000000000000
	var6 := strconv.Itoa(var5)
	pl(var6, reflect.TypeOf(var6))

	var7 := "3.14"
	if var8, err := strconv.ParseFloat(var7, 64); err == nil {
		pl(var8, err, reflect.TypeOf(var8))
	}

	var9 := fmt.Sprintf("%f", 3.14)
	pl(var9, reflect.TypeOf(var9))
}
