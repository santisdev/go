package main

import "fmt"

type User struct {
	name    string
	address string
}

func main() {
	user := User{"Charles", "Address 1"}
	p_user := &user
	p_user.name = "John"
	fmt.Println(user)
	(*p_user).name = "Charles" // golang permits us to write just p_user.name
	fmt.Println(user)
}
