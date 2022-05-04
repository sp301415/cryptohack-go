package main

import (
	"fmt"
	"math/big"
)

type Point struct {
	x           *big.Int
	y           *big.Int
	is_infinity bool
}

func NewPoint(x, y *big.Int) Point {
	if (x.Sign() == 0) && (y.Sign() == 0) {
		return Point{x: x, y: y, is_infinity: true}
	}
	return Point{x: x, y: y}
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

type Curve struct {
	a *big.Int
	b *big.Int
	p *big.Int
}

func NewCurve(a, b, p *big.Int) Curve {
	return Curve{a: a, b: b, p: p}
}

func (c Curve) Add(p, q Point) Point {
	if p.is_infinity {
		return q
	}
	if q.is_infinity {
		return p
	}

	x1 := new(big.Int).Set(p.x)
	y1 := new(big.Int).Set(p.y)
	x2 := new(big.Int).Set(q.x)
	y2 := new(big.Int).Set(q.y)
	if (x1.Cmp(x2) == 0) && (y1.CmpAbs(y2) == 0) && (y1.Sign() == -y2.Sign()) {
		return NewPoint(big.NewInt(0), big.NewInt(0))
	}

	lambda := big.NewInt(0)
	if (x1.Cmp(x2) == 0) && (y1.Cmp(y2) == 0) {
		lambda.Mul(x1, x1)
		lambda.Mul(lambda, big.NewInt(3))
		lambda.Add(lambda, c.a)
		t := new(big.Int).Mul(big.NewInt(2), y1)
		t.ModInverse(t, c.p)
		lambda.Mul(lambda, t)
		lambda.Mod(lambda, c.p)
	} else {
		lambda.Sub(y2, y1)
		t := new(big.Int).Sub(x2, x1)
		t.ModInverse(t, c.p)
		lambda.Mul(lambda, t)
		lambda.Mod(lambda, c.p)
	}
	x3 := new(big.Int).Mul(lambda, lambda)
	x3.Sub(x3, x1)
	x3.Sub(x3, x2)
	x3.Mod(x3, c.p)
	y3 := new(big.Int).Sub(x1, x3)
	y3.Mul(y3, lambda)
	y3.Sub(y3, y1)
	y3.Mod(y3, c.p)

	return NewPoint(x3, y3)
}

func (c Curve) Mul(p Point, n *big.Int) Point {
	q := NewPoint(new(big.Int).Set(p.x), new(big.Int).Set(p.y))
	r := NewPoint(big.NewInt(0), big.NewInt(0))

	for n.Sign() > 0 {
		if new(big.Int).And(n, big.NewInt(1)).Sign() > 0 {
			r = c.Add(r, q)
		}
		q = c.Add(q, q)
		n.Rsh(n, 1)
	}

	return r
}

func (c Curve) IsOnCurve(p Point) bool {
	if p.is_infinity {
		return true
	}
	x3 := new(big.Int).Set(p.x)
	x3.Mul(x3, p.x)
	x3.Mul(x3, p.x)
	ax := new(big.Int).Mul(c.a, p.x)
	yy := new(big.Int).Add(x3, ax)
	yy.Add(yy, c.b)
	yy.Mod(yy, c.p)

	pyy := new(big.Int).Mul(p.y, p.y)
	pyy.Mod(pyy, c.p)
	return yy.Cmp(pyy) == 0
}

func (c Curve) LiftX(x *big.Int) *big.Int {
	x3 := new(big.Int).Set(x)
	x3.Mul(x3, x)
	x3.Mul(x3, x)
	ax := new(big.Int).Mul(c.a, x)
	yy := new(big.Int).Add(x3, ax)
	yy.Add(yy, c.b)
	yy.ModSqrt(yy, c.p)
	return yy
}

func main() {
	E := NewCurve(big.NewInt(497), big.NewInt(1768), big.NewInt(9739))
	P := NewPoint(big.NewInt(493), big.NewInt(5564))
	Q := NewPoint(big.NewInt(1539), big.NewInt(4742))
	R := NewPoint(big.NewInt(4403), big.NewInt(5202))

	S := E.Add(P, P)
	S = E.Add(S, Q)
	S = E.Add(S, R)

	if !E.IsOnCurve(S) {
		panic("S not on Curve")
	}
	fmt.Println(S)
}
