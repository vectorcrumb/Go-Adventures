package main

import "os"
import "fmt"

func main() {
	// os.Args gives us access to raw command-line arguments.
	// The first value is the path to the current program
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	// You can get arguments with normal indexing
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
