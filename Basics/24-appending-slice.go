package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	s = append(s, 3)
	printSlice(s)

	s = append(s, 5)
	printSlice(s)

	// add more than one elemnt
	s = append(s, 2, 3)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
