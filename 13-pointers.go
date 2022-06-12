package main

import "fmt"

func main() {
	x := 8
	y := &x
	fmt.Println(x, y)
	*y = 10 //dereference and change value
	fmt.Println(x, y)
}
