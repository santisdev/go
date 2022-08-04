package main

import "fmt"

func pepe(a [5]int) bool {
	a[2] = 3
	return true
}

func main() {

	var a [5]int

	assing := func() int {
		a[0] = 5
		return 5
	}
	pepe(a)
	assing()

	fmt.Println(a)

}
