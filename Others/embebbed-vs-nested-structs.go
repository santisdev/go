package main

import "fmt"

type car struct {
	make  string
	model string
}

type embeddedTruck struct {
	car     // truck contains all the fields of the car struct
	bedSize int
}

type nestedTruck struct {
	carInfo car
	bedSize int
}

func main() {

	embTruck := embeddedTruck{
		bedSize: 10,
		car: car{
			make:  "toyota",
			model: "tacoma",
		},
	}

	nesTruck := nestedTruck{
		bedSize: 5,
		carInfo: car{
			make:  "nissan",
			model: "frontier",
		},
	}

	fmt.Println(embTruck.bedSize)

	// embedded fields promoted to the top-level
	// instead of embTruck.car.make
	fmt.Println(embTruck.make)
	fmt.Println(embTruck.model)

	fmt.Println(nesTruck.bedSize)
	fmt.Println(nesTruck.carInfo.make)
	fmt.Println(nesTruck.carInfo.model)

}
