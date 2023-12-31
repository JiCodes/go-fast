package main

import (
	"fmt"
)

// Struct and Interface

type gasEngine struct {
	mpg uint8 // miles per gallon
	gallons uint8 
	ownerInfo owner
}

type electricEngine struct {
	mpkwh uint8 // miles per kilowatt hour
	kwh uint8
}

type owner struct {
	name string
}

// assign the method to the gasEngine struct
func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

// interface
type engine interface {
	milesLeft() uint8 // method signature
}

func canMakeIt(e engine, miles uint8) {
	if e.milesLeft() >= miles{
		fmt.Println("We can make it there!")
	} else {
		fmt.Println("Uh oh, we can't make it there!")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{30, 10, owner{"Alice"}}
	myEngine.mpg = 40
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo.name)

	fmt.Printf("Total miles left: %d\n", myEngine.milesLeft())

	// anonymous struct
	var myEngine2 = struct{
		mpg uint8
		gallons uint8
	}{30, 10}
	fmt.Println(myEngine2.mpg, myEngine2.gallons)

	// interface

	var myEngine3 gasEngine = gasEngine{30, 20, owner{"Alice"}}
	canMakeIt(myEngine3, 255)

	var myEngine4 electricEngine = electricEngine{25, 15}
	canMakeIt(myEngine4, 50)

}