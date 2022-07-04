// make function; this is how you create dynamically-sized arrays.
// The make function allocates a zeroed array and returns a slice that refers to that array:
package main

import "fmt"

func main() {

	b := make([]int, 0, 5) // len(b)=0, cap(b)=5

	c := b[:cap(b)] // len(b)=5, cap(b)=5
	//since you are creating a dynamic array you can get all the elements from it,
	// and golang assings the default to each celd.

	fmt.Println(c)
	//b = b[1:] // len(b)=4, cap(b)=4
}
