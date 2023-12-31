package main

import (
	"fmt"
	"math/rand"
	"time"
)

// channel
// select statement

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 3

func main() {
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"walmart.com", "safeway.com", "costco.com"}
	for i:= range websites {
		go checkChickenPrices(websites[i], chickenChannel) // three goroutines running concurrently, checking chicken prices fluctuation
		go checkTofuPrices(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel) // waiting for the channel to send a message
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice < MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func checkTofuPrices(website string, tofuChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var tofuPrice = rand.Float32() * 20
		if tofuPrice < MAX_TOFU_PRICE {
			tofuChannel <- website
			break
		}
	}
}


func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	// var website = <-chickenChannel
	// fmt.Printf("Text sent: Found on chicken at %v\n", website)
	select {
		case website := <-chickenChannel:
			fmt.Printf("Text sent: Found on chicken at %v\n", website)
		case website := <-tofuChannel:
			fmt.Printf("Email sent: Found on tofu at %v\n", website)
	}
}