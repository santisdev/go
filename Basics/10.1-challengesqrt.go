// change the loop condition to stop once the value has stopped changing (or only changes by a very small amount
package main

import (
	"fmt"
	"math"
)

func sqrt_custom(x float64) (count int, z float64) {
	z = 1.0
	count = 0
	ep := 0.00001
	for math.Abs(z-math.Sqrt(x)) > ep {
		z -= (z*z - x) / (2 * z)
		count++
	}
	return
}

func main() {
	fmt.Println(sqrt_custom(1231))
}
