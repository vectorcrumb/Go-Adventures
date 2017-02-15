package main

import (
	"time"
	"fmt"
)

func main() {
	// select lets you wait on multiple channel operations.

	c1 := make(chan string)
	c2 := make(chan string)

	// Each channel receives a value after some time
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	// Select is used for awaiting both values simultaneously,
	// printing each one as it arrives. The for loop repeating
	// twice is to call select twice. Calling it once misses
	// the second value coming in, while calling it three or
	// more times will cause a fatal error for waiting on non
	// existent threads. Select will then only execute once,
	// but it can receive from any channel.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
