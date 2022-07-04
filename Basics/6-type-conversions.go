package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x, y int = 1, 2
	var z uint = uint(x)

	fmt.Println(reflect.TypeOf(z), reflect.TypeOf(y))
}
