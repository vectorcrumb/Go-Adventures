package main

import "fmt"

// Variadic functions take an arbitrary amount of arguments using
// the arguments... *type* syntax
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	// Call the variadic function with different arguments
	sum(1, 2)
	sum(1, 2, 3)
	// In case of having arguments gathered in a slice, you can
	// expand them with the arguments... syntax
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
