// If the concrete value inside the interface itself is nil,
// the method will be called with a nil receiver.

// In some languages this would trigger a null pointer exception,
// but in Go it is common to write methods that gracefully handle
// being called with a nil receiver (as with the method M in this example.)

// Note that an interface value that holds a nil concrete value is itself non-nil.

package main

import "fmt"

type Transport interface {
	Moving()
}

type Bmw struct {
	S string
}

func (mycar *Bmw) Moving() {
	if mycar == nil {
		fmt.Println("The Bmw does not exists")
		return
	}
	fmt.Println(mycar.S)
}

type Bike string

func (mybike Bike) Moving() {
	fmt.Println(mybike, "slow")
}

func main() {
	var i Transport
	var mybmw *Bmw
	i = mybmw
	describe(i)
	i.Moving()

	i = Bike("Bike is moving")
	describe(i)
	i.Moving()
}

func describe(i Transport) {
	fmt.Printf("(%v, %T)\n", i, i)
}
