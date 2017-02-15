package main

import "fmt"

func main() {
	// Strings can be added
	fmt.Println("go" + "lang")
	// Use at least one explicit floating point to avoid ints
	fmt.Println("1+1=", 1+1)
	fmt.Println("7/3.0=", 7/3.0)
	fmt.Println("7/3=", 7/3)
	// Boolean operators are as usual
	fmt.Println(true && false)
}
