package main

import "fmt"

func for_without_cond() {
	sum := 1
	for sum < 1000 { // this a while loop,
		sum += sum
	}
	fmt.Println(sum)
}

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	for_without_cond()
}
