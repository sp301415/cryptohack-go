package main

import "fmt"

func main() {
	s := []byte("label")
	a := 13
	for i := range s {
		s[i] ^= byte(a)
	}
	fmt.Printf("crypto{%s}\n", s)
}
