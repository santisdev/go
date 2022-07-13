package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}

func Sqrt(x float64) (float64, error) {
	var z float64 = 1.0
	ep := 0.00001
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	for math.Abs(z-math.Sqrt(x)) > ep {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
