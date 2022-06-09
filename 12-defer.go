package main

import "fmt"

// A defer statement defers the execution of a function until the surrounding function returns.
// The deferred call's arguments are evaluated immediately, but the function call is not executed
// until the surrounding function returns.

// defer calls are executed in a last in first out. That's why the first fuction its actually the last one

func main() {
	defer fmt.Println("last function")

	defer fmt.Println("world")

	fmt.Println("Hello")
}
