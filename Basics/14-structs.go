package main

import "fmt"

type Point struct {
	x int
	y int
	z int
}

func main() {
	fmt.Println(Point{1, 2, 3})
	p := Point{3, 5, 1}
	fmt.Println(p.y)
	p.z = 10
	fmt.Println(p)
}
