package main

import "fmt"

// Channels hold data, are thread safe, and listen for data when any data is added or removed from the channel

func main() {
	var c = make(chan int) // Declaration of channel with `chan` keyword
	// c <- 1                 // Adding 1 to the channel using `<-` operator
	// // We can think of a channel as containing underlying array.
	// var i = <-c // Shifting channel value into another variable.
	// // Now channel is empty.
	// fmt.Println("\nThe value of i is:", i)

	go process(c)
	fmt.Println(<-c)

}

func process(c chan int) {
	c <- 123
}
