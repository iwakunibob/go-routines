// https://youtu.be/rQQcIGqp0OU?t=724

package main

import (
	"fmt"
	"math"
	"math/rand"
)

type primemsg struct {
	num     int
	isPrime bool
}

const TEST_LEN = 100
const NUM_WORKERS = 3

// Channels are blocking: Waits for something to come down channel

func isPrime(cin chan primemsg, cout chan primemsg) { // cin=input  cout=output
	id := rand.Intn(1000000)
	i := 0
	for { // endless processing loop
		msg := <-cin // integer data comes from channel
		num := msg.num
		fmt.Println(id, "is testing", num)
		sq_root := int(math.Sqrt(float64(num)))
		for i = 2; i <= sq_root; i++ {
			if num%i == 0 {
				msg.isPrime = false
				cout <- msg
				break
			}
		}
		if i > sq_root {
			msg.isPrime = true
			cout <- msg
		}
	}
}

func main() {
	msg := primemsg{42, false}

	// Create Channels
	cin := make(chan primemsg, TEST_LEN)
	cout := make(chan primemsg, TEST_LEN)

	// Create Workers
	for i := 0; i < NUM_WORKERS; i++ {
		go isPrime(cin, cout)
	}

	// Fill the input queue
	for i := 0; i < TEST_LEN; i++ {
		msg.num = rand.Intn(1000000) + 1000000
		cin <- msg
	}

	// Read the answers
	for i := 0; i < TEST_LEN; i++ {
		msg = <-cout
		fmt.Println(msg.num, msg.isPrime)
	}
}
