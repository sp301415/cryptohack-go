package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int).SetInt64(12)
	e := new(big.Int).SetInt64(0x10001)
	n := new(big.Int).SetUint64(17 * 23)

	fmt.Println(new(big.Int).Exp(a, e, n))
}
