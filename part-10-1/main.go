package main

import (
	"fmt"
)

// Generics

func main() {
	var intSlice = []int{1, 2, 3}
	fmt.Println(sumSlice[int](intSlice))

	var float32Slice = []float32{1, 2, 3}
	fmt.Println(sumSlice[float32](float32Slice))

	var float64Slice = []float64{}
	// fmt.Println(isEmpty[float64](float64Slice))
	fmt.Println(isEmpty(float64Slice))

}

func sumSlice[T int | float32 | float64](s []T) T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}

func isEmpty[T any](s []T) bool {
	return len(s) == 0
}