package main

import (
	"fmt"
	"math"
)

// compute is a functions that recives a a parameter fn, a function with two float64 parameters and returns a float64.
// the compute function also returns a float64
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	// hypot is a function that recives floats64 x and y and returns a float64
	// we can use it like hypot(x, y)
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
