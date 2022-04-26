package main

import "fmt"

func division(dividend, divisor int) (quotient, rest int) {
	//quotient, rest = dividend/divisor, dividend-(divisor*quotient) //wrong, cant access quotient variable
	quotient = dividend / divisor
	rest = dividend - (divisor * quotient)
	return
}

func main() {
	fmt.Println(division(25, 2))
}
