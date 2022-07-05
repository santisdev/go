package main

import "fmt"

type FloatNumber float64

func (fnumber FloatNumber) Abs() float64 {
	if fnumber < 0 {
		return float64(-fnumber)
	}
	return float64(fnumber)
}

func main() {
	f := FloatNumber(-5)
	fmt.Println(f.Abs())
}
