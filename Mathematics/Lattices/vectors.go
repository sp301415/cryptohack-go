package main

import (
	"fmt"
)

type Vec3 [3]int

func (v Vec3) Mul(x int) Vec3 {
	r := Vec3{}
	for i := range v {
		r[i] = x * v[i]
	}
	return r
}

func (v Vec3) Sub(w Vec3) Vec3 {
	r := Vec3{}
	for i := range v {
		r[i] = v[i] - w[i]
	}
	return r
}

func (v Vec3) Dot(w Vec3) int {
	r := 0
	for i := range v {
		r += v[i] * w[i]
	}
	return r
}

func (v Vec3) String() string {
	return fmt.Sprintf(`(%d, %d, %d)`, v[0], v[1], v[2])
}

func main() {
	v := Vec3{2, 6, 3}
	w := Vec3{1, 0, 0}
	u := Vec3{7, 7, 2}

	r := ((v.Mul(2).Sub(w)).Mul(3)).Sub(u.Mul(2))
	fmt.Println(r)
}
