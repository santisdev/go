package main

import "fmt"

func changeValue(str *string) {
	*str = "Changed!"
}

func changeValue2(str string) {
	str = "change it again"
}

func main() {
	toChange := "Text to change"
	fmt.Println(toChange)
	changeValue(&toChange)
	fmt.Println(toChange)

	fmt.Println("Current value:", toChange)
	changeValue2(toChange) //doesn't change the value
	fmt.Println("Current value:", toChange)
}
