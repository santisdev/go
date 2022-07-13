// A type assertion provides access to an interface value's underlying concrete value.

// t := i.(T)

// This statement asserts that the interface value i holds the concrete type T and
// assigns the underlying T value to the variable t.

// If i does not hold a T, the statement will trigger a panic.

// To test whether an interface value holds a specific type, a type assertion can
// return two values: the underlying value and a boolean value that reports whether
// the assertion succeeded.

// t, ok := i.(T)

package main

import (
	"fmt"
)

func main() {
	var i interface{} = "hello"

	p := i.(string)
	fmt.Println(p)

	q, ok := i.(string)
	fmt.Println(q, ok)

	r, ok := i.(float64)
	fmt.Println(r, ok)

	s := i.(float64) // panic
	fmt.Println(s)
}
