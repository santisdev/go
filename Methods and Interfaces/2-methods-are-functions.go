package main

import "fmt"

const kph_mph_cf = 1.609344

type Car struct {
	speed, km float64
}

func SpeedInMiles(c *Car) {
	c.speed = c.speed / kph_mph_cf
}

func main() {
	first_car := Car{60, 234}
	fmt.Println(first_car)
	SpeedInMiles(&first_car)
	fmt.Println(first_car)
}
