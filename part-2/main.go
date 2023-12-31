package main

import (
	"fmt" 
	"unicode/utf8"
)

// Data Types and constant variables

func main() {
	var intNum int // default value is 0
	fmt.Println(intNum)

	var floatNum float64 = 12345.6789
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 * float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	var myString string = "Hello, World!"
	myString2 := "Hello, World!" // short variable declaration
	fmt.Println(myString2)

	var myString1 string = `Hello,
	World!`

	fmt.Println(myString)
	fmt.Println(len(myString1)) // returns bytes 15 (including the newline character)

	fmt.Println(utf8.RuneCountInString("γ")) // returns 1
	var myRune rune = 'a'
	fmt.Println(myRune) // returns 97

	var myBool bool // default value is false
	fmt.Println(myBool)


	var var1, var2 = 1, 2
	fmt.Println(var1, var2)
	var3, var4 := 3, 4
	fmt.Println(var3, var4)

	var myVar string = foo()
	fmt.Println(myVar)

	const myConst string = "Hello, World!" // constant cannot be changed
	fmt.Println(myConst)

	const pi float64 = 3.14159265359
}

func foo() string {
	return "Hello, foo!"
}