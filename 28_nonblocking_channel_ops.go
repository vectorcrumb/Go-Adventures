package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan string)

	// Using the default case, we can implement a non-blocking channel receive
	select {
	case msg := <-messages:
		fmt.Println("received message:", msg)
	default:
		fmt.Println("no message received")
	}

	// In the same way as before, we've implemented a non-blocking channel send
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message:", msg)
	default:
		fmt.Println("no message sent")
	}

	// Finally, we can implement a non-blocking multi-way select block.
	select {
	case msg := <-messages:
		fmt.Println("received message:", msg)
	case sig := <-signals:
		fmt.Println("received signal:", sig)
	default:
		fmt.Println("no activity")
	}
}
