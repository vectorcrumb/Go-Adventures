package main

import "fmt"

// This function returns another function which returns an int. It
// is an anonymous function which closes over i to form a closure.
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	// Assign the inner function to nextInt and increment i
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	// The state is unique to every instance
	newInts := intSeq()
	fmt.Println(newInts())
}
