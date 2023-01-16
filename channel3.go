package main

import "fmt"

// Channels are blocking: Waits for something to come down channel

func sum(c chan int, co chan int) { // c=input  co=output
	for { // endless processing loop
		x := <-c // integer data comes from channel to x
		y := <-c // integer data comes from channel to y
		fmt.Println(x, "+", y, "=", x+y)
		co <- x + y // sum of x + y enters channel
	}
}

func main() {
	c := make(chan int, 6)  // rcv 6 values from input channel
	co := make(chan int, 3) // snd 3 values to ouput channel
	go sum(c, co)
	c <- 10 // 10 enters input channel
	c <- 15
	c <- 99
	c <- 1
	c <- 23
	c <- 7

	result := <-co // result exits channel
	fmt.Println(result)
	result = <-co // result exits channel
	fmt.Println(result)
	result = <-co // result exits channel
	fmt.Println(result)
}
