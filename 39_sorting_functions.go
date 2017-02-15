package main

import "sort"
import "fmt"

// We start of by creating a custom type, which, in this case, is simply an alias
// to []string. The alias must indicate which data type we'll be sorting and the
// name will be the function to call inside sort.Sort()
type ByLength []string

// Then we implement the 3 functions from sort.Interface: Len, Less and Swap. This
// will allow us to call the generic version of Sort. Less holds the actual sorting logic
func (s ByLength) Len() int{
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}