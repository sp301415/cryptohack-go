package main

import (
	"fmt"
	"math/big"
)

func main() {
	e := big.NewInt(65537)
	p, _ := new(big.Int).SetString("857504083339712752489993810777", 10)
	q, _ := new(big.Int).SetString("1029224947942998075080348647219", 10)
	N := new(big.Int).Mul(p, q)

	p.Sub(p, big.NewInt(1))
	q.Sub(q, big.NewInt(1))
	toient := new(big.Int).Mul(p, q)

	d := new(big.Int).ModInverse(e, toient)

	c, _ := new(big.Int).SetString("77578995801157823671636298847186723593814843845525223303932", 10)
	fmt.Println(c.Exp(c, d, N))
}
