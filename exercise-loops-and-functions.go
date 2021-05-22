package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	const accuracy = 1000000000
	z := float64(1)

	for {
		c := z
		z -= (z*z - x) / (2 * z)
		if int(c*accuracy) == int(z*accuracy) {
			return z
		}
	}
}

func main() {
	num := 28524.0
	fmt.Println(Sqrt(num))
	fmt.Println(math.Sqrt(num))
}
