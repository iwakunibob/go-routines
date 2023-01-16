package main

import "fmt"

// Channels are blocking: Waits for something to come down channel

func sum(c chan int) {
	x := <-c   // integer data comes from channel to x
	y := <-c   // integer data comes from channel to y
	c <- x + y // sum of x + y enters channel
}

func main() {
	c := make(chan int, 2) // get two values from channel ERROR
	go sum(c)
	c <- 10       // 10 enters channel
	c <- 15       // 15 enters channel
	result := <-c // result exits channel
	fmt.Println(result)
}
