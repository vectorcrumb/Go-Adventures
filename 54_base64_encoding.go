package main

import b64 "encoding/base64"
import "fmt"

func main() {

	data := "abc123!?$*&()'-=@~"
	// EncodeToString we can convert to standard base 64. We use []byte() to cast to a slice of bytes
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)
	// DecodeString takes a base64 string and returns the original data and an error
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	// We use the same functions, different package for URL compatible base64
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}