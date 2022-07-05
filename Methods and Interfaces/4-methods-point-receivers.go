package main

import "fmt"

const kph_mph_cf = 1.609344

type Car struct {
	speed, km float64
}

func (c *Car) SpeedInMiles() {
	c.speed = c.speed / kph_mph_cf
}

// A method is just a function with a receiver argument.
func main() {
	first_car := Car{60, 234}
	fmt.Println(first_car)
	first_car.SpeedInMiles()
	//notice this is not a pointer, its a golang convenience and first_car.method is (&first_car).method
	// since the SpeedInMiles method has a pointer receiver.
	fmt.Println(first_car)
}

// more details  -----

// var v Vertex
// ScaleFunc(v, 5)  // Compile error!
// ScaleFunc(&v, 5) // OK

// while methods with pointer receivers take either a value or a pointer as the receiver when they are called:

// var v Vertex
// v.Scale(5)  // OK
// p := &v
// p.Scale(10) // OK

// Why use pointer receivers?
// The first reason is so that the method can modify the value that its receiver points to.

// The second is to avoid copying the value on each method call.
// This can be more efficient if the receiver is a large struct, for example.
