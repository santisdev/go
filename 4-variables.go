package main

import (
	"fmt"
	"reflect"
)

var c, python, java bool //defaults to false

func main() {
	var i int //defaults to 0
	// var j, k int = 1, 2
	fmt.Println(i, c, python, java)

	var c, python, java = true, false, "no!" // variable will take the type of the initializer
	fmt.Println(reflect.TypeOf(c), reflect.TypeOf(python), reflect.TypeOf(java))

}
