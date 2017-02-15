package main

import "fmt"
import "sort"

func main() {
	// Sorts are specific to buildin types. Sorting occurs in place
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)
	// We can also sort ints
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)
	// Finally, we can check if a slice is already sorted
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)
}
