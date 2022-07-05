package main

import "fmt"

type Transport interface {
	Moving()
}

type Bmw struct {
	S string
}

func (mycar *Bmw) Moving() {
	fmt.Println(mycar.S)
}

type Bike string

func (mybike Bike) Moving() {
	fmt.Println(mybike, "slow")
}

func main() {
	var i Transport
	i = &Bmw{"Car is moving"}
	describe(i)
	i.Moving()

	i = Bike("Bike is moving")
	describe(i)
	i.Moving()
}

func describe(i Transport) {
	fmt.Printf("(%v, %T)\n", i, i)
}
