package main

import (
	"fmt"
)

type ErrNavigateSqrt float64

func (e ErrNavigateSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNavigateSqrt(x)
	}

	const accuracy = 1000000000
	z := float64(1)

	for {
		c := z
		z -= (z*z - x) / (2 * z)
		if int(c*accuracy) == int(z*accuracy) {
			return z, nil
		}
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
