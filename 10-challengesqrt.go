package main

import (
	"fmt"
)

func sqrt_custom(x float64) (i int, z float64) {
	z = 1.0
	for i = 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return
}

func main() {
	fmt.Println(sqrt_custom(100))
}
