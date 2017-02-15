package main

import "fmt"

func main() {
	// Create a map with make(map[key-type]value-type)
	m := make(map[string]int)
	// Set pairs using map[key] = val
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)
	// Get values using map[key[
	v1 := m["k1"]
	fmt.Println("v1:", v1)
	// Call len() for number of pairs
	fmt.Println("len:", len(m))
	// Call delete(map, key) to remove a pair
	delete(m, "k2")
	fmt.Println("map:", m)
	// Receive a second value to check if the key pair existed
	// _ is a blank identifier (the compiler won't complain if
	// it isn't being used)
	_, prs := m["k2"]
	fmt.Println("prs:", prs)
	// A map may be declared and initialized in the same line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
