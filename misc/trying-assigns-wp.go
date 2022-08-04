package main

import "fmt"

func assign1(arr *[]*int) {
	b := 3
	*arr = append(*arr, &b)
}

func main() {

	a := []*int{}

	assing2 := func() {
		c := 5
		a = append(a, &c)
	}
	assign1(&a)
	assing2()
	assing2()

	for index, n := range a {
		fmt.Printf("Element at index %d => %d \n", index, *n)
	}

}
