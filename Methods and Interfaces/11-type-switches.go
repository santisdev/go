// A type switch is a construct that permits several type assertions in series.

// A type switch is like a regular switch statement, but the cases in a type switch
// specify types (not values), and those values are compared against the type of the
// value held by the given interface value.

// The declaration in a type switch has the same syntax as a type assertion i.(T),
// but the specific type T is replaced with the keyword type.

// In the default case (where there is no match), the variable v is of the same interface
// type and value as i.

package main

import (
	"fmt"
)

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}

}

func main() {
	do(21)
	do("hello")
	do(true)
}