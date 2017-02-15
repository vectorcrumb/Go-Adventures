package main

import "fmt"

func main() {

	// Make a slice of string, length 3
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// Write to elements of the slice as an array
	s[0] = "a"
	s[1], s[2] = "b", "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	// Get the length with len(), capacity with cap()
	fmt.Println("len:", len(s))

	// Append elements to the array. Grab the return value!
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Make a new empty slice and copy another into it
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slice operations akin to Python. This obtains positions 2, 3 and 4
	l := s[2:5]
	fmt.Println("sl1", l)
	// Empty values default to 0 or len()
	l = s[:5]
	fmt.Println("sl2", l)
	l = s[2:]
	fmt.Println("sl3", l)

	// Declare and initialize a slice
	t := []string{"g", "h", "i"}
	fmt.Println("dcl", t)

	// Slices can be multidimensional
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

}
