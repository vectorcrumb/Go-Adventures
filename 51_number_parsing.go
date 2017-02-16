package main

import "strconv"
import "fmt"

func main() {

	// With ParseFloat, we specify how many bits of
	// precision to parse (in this case, 64 bits)
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// With ParseInt, we use 0 to make it infer the
	// base and 64 to determine the variable size
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// ParseInt will recognize hex-formatted numbers
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// ParseUint is also available
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// Atoi is a simple function for base-10 int parsing
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// Parse functions return an error on bad input
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
