package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// Closure is an inner function that has access to all the variables in the scope where it was created.
func main() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq() // if you create a new one, you lose the previous state
	fmt.Println(newInts())
}
