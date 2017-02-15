package main

import "fmt"

func main() {
	// Range can be used to iterate over the elements in a slice
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum", sum)

	// range returns two values: index, element
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index", i)
		}
	}
	// range can return key value pairs of a map
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	// However, it can also just return the keys of a map
	for k := range kvs {
		fmt.Println("key", k)
	}
	// On a string, range iterates over Unicode code points
	// The first value is the index of the rune and the second
	// the code point or rune itself
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
