package main

import "fmt"

func main() {
	x := 8
	y := &x
	fmt.Println(x, y)
	*y = 10 //dereference
	fmt.Println(x, y)
}
