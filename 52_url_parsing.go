package main

import (
	"fmt"
	"net"
	"net/url"
)

func main () {
	// This example URL includes a scheme, auth info, host,
	// port, path, query params and query fragment
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	// Parse the URL and ensure there are no errors
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	// Accessing the scheme is simple
	fmt.Println(u.Scheme)
	// User contains all auth info
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)
	// Host contains connection info. Use SplitHostPort to
	// extract the hostname and port
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)
	// Now we extract the path and the fragmet
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	// To get query strings in a k=v string format, use RawQuery.
	// You can also parse query params into a map. The parsed info
	// becomes slices of strings, so we use [0] to access it
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])

}
