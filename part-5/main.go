package main

import (
	"fmt"
	"strings"
)

// Strings, Runes, and Bytes

// UTF-8 is how Go represents strings internally

func main() {
	// var myString string = "résumé"
	var myString = []rune("résumé") // rune is an alias for int32
	var indexed = myString[1]
	fmt.Printf("indexed is %v %T\n", indexed, indexed) 

	for i, v := range myString {
		fmt.Printf("index is %v and value is %v\n", i, v)
	}
	fmt.Printf("length is %v\n", len(myString)) 

	var myRune = 'a' // single quotes for runes
	fmt.Printf("myRune is %v %T\n", myRune, myRune)

	var strSlice = []string{"c", "a", "t"}
	var catStr = ""
	for i := range strSlice {
		catStr += strSlice[i] // string concatenation is inefficient
	}
	fmt.Printf("catStr is %v\n", catStr)
	// catStr[0] = 'b' // cannot assign to catStr[0] once created


	// instead use strings.Builder after importing "strings"
	// the array is allocated internally and values are appended when calling WriteString method
	// only at the end is the new string created from appended values
	var strBuilder strings.Builder

	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	catStr = strBuilder.String()
	fmt.Printf("catStr is %v\n", catStr)

}