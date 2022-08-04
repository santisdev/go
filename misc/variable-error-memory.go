package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Declare a list/slice of string pointers
	listOfNumberStrings := []*string{}

	// Forward declaration of the variable we will store the data in before adding to slice
	var numberString string

	// Loop from 0 to 9
	for i := 0; i < 10; i++ {
		// Format the string with '#' prefixed to the number
		numberString = fmt.Sprintf("#%s", strconv.Itoa(i)) //this is an error. Re-declare the variable numberString so it gets a new unique memory address.
		fmt.Printf("Adding number %s to the slice\n", numberString)
		// Add number string to the slice
		listOfNumberStrings = append(listOfNumberStrings, &numberString)
	}

	for _, n := range listOfNumberStrings {
		fmt.Printf("%s\n", *n)
	}

	fixed()
}

func fixed() {
	// Declare a list/slice of string pointers
	listOfNumberStrings := []*string{}

	// Loop from 0 to 9
	for i := 0; i < 10; i++ {
		var numberString string
		// Format the string with '#' prefixed to the number
		numberString = fmt.Sprintf("#%s", strconv.Itoa(i)) //this is an error. Re-declare the variable numberString so it gets a new unique memory address.
		fmt.Printf("Adding number %s to the slice\n", numberString)
		// Add number string to the slice
		listOfNumberStrings = append(listOfNumberStrings, &numberString)
	}

	for _, n := range listOfNumberStrings {
		fmt.Printf("%s\n", *n)
	}
}
