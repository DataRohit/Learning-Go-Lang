package main

import (
	"fmt"
	"math"
)

var pl = fmt.Println

func main() {
	pl(math.Abs(-10))
	pl(math.Pow(4, 2))
	pl(math.Sqrt(16))
	pl(math.Cbrt(8))
	pl(math.Ceil(4.4))
	pl(math.Floor(4.4))
	pl(math.Round(4.4))
	pl(math.Log2(8))
	pl(math.Log10(100))
	pl(math.Log(math.Exp(2)))
	pl(math.Max(5, 4))
	pl(math.Min(5, 4))

	r90 := 90 * math.Pi / 180
	pl(r90)

	d90 := r90 * 180 / math.Pi
	pl(d90)

	pl(math.Sin(r90))
}
