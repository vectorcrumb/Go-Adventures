package main

import "fmt"

// the (int, int) indicates the functions returns 2 ints
func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	// Return values can be discarded with the blank identifier
	_, c := vals()
	fmt.Println(c)
}
