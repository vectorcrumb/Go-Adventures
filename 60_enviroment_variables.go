package main

import (
	"os"
	"strings"
	"fmt"
)

func main() {
	// Set environment variables like this
	os.Setenv("FOO", "1")
	// Get environment variables like this. If the variable
	// doesn't exist, an empty string is returned
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))
	// We can also obtain all key-value pairs of environment variables
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}