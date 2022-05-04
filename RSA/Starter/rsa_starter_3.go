package main

import (
	"fmt"
	"math/big"
)

func main() {
	p, _ := new(big.Int).SetString("857504083339712752489993810777", 10)
	q, _ := new(big.Int).SetString("1029224947942998075080348647219", 10)

	p.Sub(p, big.NewInt(1))
	q.Sub(q, big.NewInt(1))

	fmt.Println(p.Mul(p, q))
}
