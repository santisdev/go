package main

import "fmt"

const kph_mph_cf = 1.609344

type Car struct {
	speed, km float64
}

func (c *Car) SpeedInMiles() {
	c.speed = c.speed / kph_mph_cf
}

// A method is just a function with a receiver argument.
func main() {
	first_car := Car{60, 234}
	fmt.Println(first_car)
	first_car.SpeedInMiles()
	fmt.Println(first_car)
}
