package main

import (
	"fmt"
)

// generic used with struct types
type gasEngine struct {
	gallons float32
	mpg		 float32
}

type electricEngine struct {
	kwh float32
	mpkwh float32
}

type car[T gasEngine | electricEngine] struct {
	carMake string
	carModel string
	engine T
}

func main() {
	var gasCar = car[gasEngine]{
		carMake: "Ford",
		carModel: "F150",
		engine: gasEngine{
			gallons: 10,
			mpg: 20,
		},
	}
	var electricCar = car[electricEngine]{
		carMake: "Tesla",
		carModel: "Model S",
		engine: electricEngine{
			kwh: 85,
			mpkwh: 3.5,
		},
	}

	fmt.Println(gasCar)
	fmt.Println(electricCar)
}