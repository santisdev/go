package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2] // s[0], s[1]
	fmt.Println(s)

	s = s[1:] // from s[1] to the end (one element in this case)
	fmt.Println(s)
}
