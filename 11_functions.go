package main

import "fmt"

// Unlike JS, an explicit return statement is required
func plus(a int, b int) int {
	return a + b
}

// comma separated values all share the same type
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2=", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3=", res)
}