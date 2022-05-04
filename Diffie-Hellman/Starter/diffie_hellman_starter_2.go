package main

import (
	"fmt"
	"math/big"
)

func main() {
	one := big.NewInt(1)
	p := big.NewInt(28151)

	for k := big.NewInt(1); k.Cmp(p) < 0; k.Add(k, one) {
		is_gen := true
		for i := big.NewInt(2); i.Cmp(p) < 0; i.Add(i, one) {
			x := new(big.Int).Exp(k, i, p)
			if x.Cmp(k) == 0 {
				is_gen = false
				break
			}
		}

		if is_gen {
			fmt.Println(k)
			break
		}
	}
}
