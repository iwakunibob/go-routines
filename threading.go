// This is a demo of Go routins by Gary Explains on Youtube
// https://youtu.be/rQQcIGqp0OU

package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("Hello") // to create a thread insert go keyword
	say("Earth")
}
