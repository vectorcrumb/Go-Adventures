package main

import "strings"
import "fmt"


// Return the first index of the target string or -1 if no match is found
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Returns true if the target string t is in the slice
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// Returns true if a string in the slice satisfies the predicate f
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// Returns true if all strings in the slice satisfy the predicate f
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Return a new slice with the results of applying f to each string in the original slice
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// Go doesn't include support for generics. It's common to provide collection functions
// if and when they're necessary for your program and data types. This file includes some
// examples of collection functions for slices of strings.

func main () {
	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))
	fmt.Println(Include(strs, "grape"))
	fmt.Println(Any(strs, func (v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(All(strs, func (v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))
	fmt.Println(Map(strs, strings.ToUpper))
}

