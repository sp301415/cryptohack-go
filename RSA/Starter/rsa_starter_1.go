package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int).SetInt64(101)
	b := new(big.Int).SetInt64(17)
	n := new(big.Int).SetUint64(22663)

	fmt.Println(new(big.Int).Exp(a, b, n))
}
