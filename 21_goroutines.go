package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// This is a normal, synchronous function call to f(s)
	f("direct")
	// By using the keyword 'go', we invoke this function in a goroutine.
	// This new goroutine executes concurrently with the calling one.
	go f("goroutine")
	// goroutines may also be started for anonymous functions
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	// The two previous function calls are running async, so
	// the code execution falls all the way to the Scanln line

	// Scanln takes an input from stdin and writes it to the memaddr
	// of input. It's simply used to pause the program in this case
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}