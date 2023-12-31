package main

import (
	"errors"
	"fmt"
)

// Functions and if / switch statements
func main() {
	var printValue string = "Hello World!"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 2
	var result, remainder, err = intDivision(numerator, denominator)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else if remainder == 0 {
	// 	fmt.Printf("the result of integer division is %v\n", result)
	// } else {
	// fmt.Printf("the result of integer division is %v with remainder %v\n", result, remainder)
	// }

	switch {
		case err != nil:
			fmt.Println(err.Error())
			// break not needed in go
		case remainder == 0:
			fmt.Printf("the result of integer division is %v\n", result)
		default:
			fmt.Printf("the result of integer division is %v with remainder %v\n", result, remainder)
	}

	switch remainder {
		case 0:
			fmt.Println("The division was exact")
		case 1, 2:
			fmt.Println("The division was close")
		default:
			fmt.Println("The division was not close")
		}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		// err = fmt.Errorf("cannot divide by zero")
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}