package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] //primes from primes[1] to primes[4]
	fmt.Println(s)

	// slices includes the first one, but excludes the last one  [low: high)

	// a slice describes a section of underlying data, this means that changing the elements of
	// the slice, modifies the corresponding elements of its underlying array,
	// and other slices that share the same underlying array will see those changes too

	s[0] = 19 // this changes primes[1]

	fmt.Println(primes)
}
