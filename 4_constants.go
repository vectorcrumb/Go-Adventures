package main

import "fmt"
import "math"

const s string = "constant"

func main() {
	// const cannot change value
	fmt.Println(s)
	// Type can be inferred
	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)
	// But not specifically. In this case, an explicit cast to int64
	// is required to satisfy the argument type in math.Sin()
	fmt.Println(int64(d))

	fmt.Println(math.Sin(d))
}
