package main

import (
	"fmt"
	"os"
)
//import "os"

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}

	// The %v verb prints a struct
	fmt.Printf("%v\n", p)
	// The %+v variant will print the struct fields
	fmt.Printf("%+v\n", p)
	// The %#v variant prints a source code representation
	fmt.Printf("%#v\n", p)
	// To print a value's type, use %T
	fmt.Printf("%T\n", p)

	// To format a boolean, use %t
	fmt.Printf("%t\n", true)
	// For printing base-10 integers, we use %d
	fmt.Printf("%d\n", 123)
	// For a binary representation of a int, use %b
	fmt.Printf("%b\n", 14)
	// For a given int, we can print the character it represents using %c
	fmt.Printf("%c\n", 33)
	// %x provides hex encoding
	fmt.Printf("%x\n", 456)
	// For basic decimal formatting for floats, we use %f
	fmt.Printf("%f\n", 78.9)
	// %e and %E format in different variations of scientific notation
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	// For basic string printing, use %s
	fmt.Printf("%s\n", "\"string\"")
	// For double quoting strings in Go source, use %q
	fmt.Printf("%q\n", "\"string\"")
	// To render strings in hex, as with integers, use %x
	fmt.Printf("%x\n", "hex this")
	// To print a representation of a pointer, use %p
	fmt.Printf("%p\n", &p)

	// To control the width of an int, place a number before the flag
	fmt.Printf("|%6d|%6d|\n", 12, 345)
	// For a float, place a .n after the columns
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	// To left justify, use a - sign
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	// The same for strings
	fmt.Printf("|%6s|%6s|\n", "foo", "b")
	// And for left justification
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)
	fmt.Fprintf(os.Stderr, "an %s\n", "error")

}