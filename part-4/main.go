package main

import (
	"fmt"
)

// Arrays, Slices, and Maps
// Loops control structures

// []Arrays fixed length, same type, indexable, contiguous memory

func main() {
	var intArray [3]int32 = [3]int32{1, 2, 3} // 12 bytes
	// intArray := [...]int32{1, 2, 3} 
	fmt.Println(intArray)
	
	intArray[1] = 111
	fmt.Println(intArray[0])
	fmt.Println(intArray[1:3]) 

	fmt.Println(&intArray[0]) // pointer / address of first element 
	fmt.Println(&intArray[1]) // contiguous memory location
	fmt.Println(&intArray[2]) 


	var intSlice []int32 = []int32{4, 5, 6} 
	fmt.Printf("the length is %v and the capacity is %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("the length is %v and the capacity is %v\n", len(intSlice), cap(intSlice))
	// [4, 5, 6, 7, *, *] 

	// In Go, slices are dynamic, meaning they can grow as you append elements.
	// When the underlying array of the slice is full and you append an element, 
	// Go creates a new array with double the capacity of the previous array and copies the elements to this new array.

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...) // append multiple elements using spread operator 
	fmt.Printf("the length is %v and the capacity is %v\n", len(intSlice), cap(intSlice))

	var intSlice3 []int32 = make([]int32, 3, 8) // make a slice with length 3 and capacity 8
	fmt.Printf("the length is %v and the capacity is %v\n", len(intSlice3), cap(intSlice3))


	// iterate over a array / slice
	for i, value := range intSlice {
		fmt.Printf("index is %v and value is %v\n", i, value)
	}


	// Maps 
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)
	fmt.Println(myMap["Adam"]) // 0

	var myMap2 map[string]int32 = map[string]int32{"Adam": 16, "Bob": 20, "Cindy": 35}
	var age, ok = myMap2["Adam"]
	if ok {
		fmt.Printf("age is %v\n", age)
	} else {
		fmt.Println("Adam not found")
	}

	delete(myMap2, "Adam")

	// iterate over a map, order is not reserved
	for key, value := range myMap2 {
		fmt.Printf("key is %v and value is %v\n", key, value)
	}

}