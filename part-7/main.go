package main

import (
	"fmt"
)

// Pointer
// special type of variable that holds the memory address of another variable
// it take 32 bits on a 32 bit system and 64 bits on a 64 bit system

func main() {
	var p *int32 = new(int32) // new() allocates memory for the type and returns a pointer
	var p1 *int32 
	// this is a nil pointer, 
	// we didn't allocate memory for it 
	// so we can't assign a value to it
	var i int32 = 10


	fmt.Printf("The value of p is: %v\n", p)
	fmt.Printf("The value p points to is: %d\n", *p) // * - dereference operator
	
	fmt.Printf("The value of i is: %v\n", i)
	fmt.Printf("The address of i is: %v\n", &i) // & - address of operator

	p = &i // assign the address of i to p
	*p = 20 // assign 20 to the memory address that p points to
	fmt.Printf("The value p points to is now: %d\n", *p)
	fmt.Printf("The address of i is also changed to: %v\n", i)

	fmt.Printf("The value of p1 is: %v\n", p1)





	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("The memory location of thing1 array is: %p\n", &thing1)
	var result [5]float64 = square(&thing1)
	fmt.Printf("The result is: %v\n", result)
	fmt.Printf("The value of thing1 is: %v\n", thing1)

}

// make a function that takes a pointer to an array and squares each element
// this will save us memory because we don't have to copy the array
func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("The memory location of thing2 array is: %p\n", thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}