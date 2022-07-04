package main

import (
	"fmt"
	"reflect"
)

func main() {
	const pi = 3.14
	fmt.Println(reflect.TypeOf(pi))
	fmt.Println(pi)
}
