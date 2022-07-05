// An interface type is defined as a set of method signatures.

// A value of interface type can hold any value that implements those methods.

// A type implements an interface by implementing its methods.
// There is no explicit declaration of intent, no "implements" keyword

package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	a = v // To fix it chage it to a = &v or delete de * in *Vertex in line 44

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 { // MyFloat implements the interface Abser
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 { // *Vertex implements the interface Abser
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
