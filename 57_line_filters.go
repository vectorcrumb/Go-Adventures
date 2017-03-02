package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// We can wrap the unbuffered Stdin with a buffered scanner
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// Text() returns the current token (the whole line)
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr, "error:", err)
		os.Exit(1)
	}
}