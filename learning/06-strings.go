package main

import (
	"fmt"
	"strings"
)

var pl = fmt.Println

func main() {
	str1 := "Strings"
	replacer := strings.NewReplacer("S", "So")
	str2 := replacer.Replace(str1)
	pl(str1, str2)
	pl(len(str1), len(str2))
	pl(strings.Contains(str2, "ring"))
	pl(strings.Index(str2, "o"))
	pl(strings.Replace(str2, "o", "0", -1))

	str3 := "\nEscape Characters\n"
	pl(str3)

	str3 = strings.TrimSpace(str3)
	pl(str3)

	pl(strings.Split("a-b-c-d", "-"))

	pl(strings.ToLower(str3))
	pl(strings.ToUpper(str3))

	pl(strings.HasPrefix(str1, "Str"))
	pl(strings.HasSuffix(str1, "gs"))
}
