package main

import (
	"time"
	"fmt"
	"math/rand"
)

func main() {
	p, pln := fmt.Print, fmt.Println
	// n := Intn(N) returns a random int given 0 <= n < N
	p(rand.Intn(100), ", ")
	p(rand.Intn(100))
	pln()
	// rand.Float64() returns a f such that 0.0 <= f < 1.0
	pln(rand.Float64())
	// We can use this this property to expand the rang of
	// Float64 with linear functions
	p((rand.Float64() * 5) + 5, ", ")
	p((rand.Float64() * 5) + 5)
	pln()
	// rand is based on a PRNG, so it's deterministic. Give
	// it a seed to change the values it generates.
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	// Call the resulting rand.Rand just as before with rand
	p(r1.Intn(100), ",")
	p(r1.Intn(100))
	pln()
	// Seeding a source with the same number generates the
	// same sequence of random numbers
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}
