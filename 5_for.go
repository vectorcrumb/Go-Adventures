package main

import "fmt"

func main() {
	// for can take only a condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}
	// Or it can be used as in C/C++ and Java
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	// There is no while operator. Use an empty for with break to escape it
	for {
		fmt.Println("loop")
		break
	}
	// for also accepts the continue statement, which breaks out of the
	// current cycle and resumes normally on the next one
	for n := 0; n <= 5; n++ {
		if n % 2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
