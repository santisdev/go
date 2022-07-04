package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

// range on arrays and slices provides both the index and value for each entry
// you can skip the index (i) or value(v) by assigning it to _
// for i, _ := range pow
// or for _, v := range pow
