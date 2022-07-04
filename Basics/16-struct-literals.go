package main

import "fmt"

type User struct {
	name, address string
}

var (
	v1 = User{"Charles", "Charles Address"} //has type User
	v2 = User{name: "John"}                 //address = ''
	v3 = User{}                             // name and address = ''
	p  = &User{"John", "John Address"}      // has type *User
)

func main() {
	fmt.Println(v1, v2, v3, p)
}
