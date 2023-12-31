package main

import (
	"fmt"
)

// channel is a way to communicate between goroutines

// hold data
// Thread safe
// listen for data

func main() {
	var c = make(chan int)
	go process(c)
	// var i = <-c // retrieve data from channel
	
	for i := range c{
		fmt.Println(i)
	}
}

func process(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i // add data to channel
	}
	close(c)
}