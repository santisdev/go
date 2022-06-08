package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i, j int = 1, 2
	k := 3
	flag := false
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java, flag)

	fmt.Println(reflect.TypeOf(i), reflect.TypeOf(j), reflect.TypeOf(k), reflect.TypeOf(c),
		reflect.TypeOf(python), reflect.TypeOf(java), reflect.TypeOf(flag))
}
