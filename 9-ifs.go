package main

import (
	"fmt"
	"math"
)

func eval_string(text string) bool {
	if text == "This is true" {
		return true
	} else {
		return false
	}
}

func pow(x, n, limit float64) float64 { //statement before condition, v is in the scope until the en of the if
	if v := math.Pow(x, n); v < limit {
		return v
	}
	return limit
}

func main() {
	fmt.Println(eval_string("This is true"), eval_string("Hello"))

	fmt.Println(pow(3, 3, 100))
	fmt.Println(pow(3, 3, 20))
}
