package main

import "fmt"

// fmt.Println([]byte("café")) // [99 97 102 195 169]
// fmt.Println([]rune("café")) // [99 97 102 233]

func reverse(s string) string {
	rns := []rune(s) //slice
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func main() {

	var name string

	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Hello", name)

	//reverse
	fmt.Println("The reverse is:", reverse(name))
}
