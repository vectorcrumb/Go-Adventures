package main

import "fmt"

func main() {
	// Syntax: var *name* *type* = *value*
	var a string = "initial"
	fmt.Println(a)
	// Comma separated variables share the same type
	var b, c int = 1, 2
	fmt.Println(b, c)
	// The compiler can infer variable type
	var d = true
	fmt.Println(d)
	// Declaration starts variable at a zero value
	var e int
	fmt.Println(e)
	e = 1
	fmt.Println(e)
	// Shorthand declaration and initialization
	f := "short"
	fmt.Println(f)
}
