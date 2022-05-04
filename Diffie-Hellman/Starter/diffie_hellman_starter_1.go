package main

import (
	"fmt"
	"math/big"
)

func main() {
	p := big.NewInt(991)
	g := big.NewInt(209)
	fmt.Println(g.ModInverse(g, p))
}
