package main

import (
	"fmt"
)

const Pi = 3.14159265358979323846

const (
	A = iota
	B
	C
)

type Dia int

const (
	Domingo Dia = iota
	Segunda
	Terca
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
)

func main() {
	fmt.Println("---------")
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println("---------")
	fmt.Println(Domingo)
	fmt.Println(Segunda)
	fmt.Println(Terca)
	fmt.Println("---------")
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)

	var f32 float32 = Pi
	var f64 float64 = Pi
	fmt.Println(f32)
	fmt.Println(f64)

	const (
		a, b = iota, iota + 10
		c, d
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	const big = 1 << 100
	fmt.Println(big >> 98)

	// fmt.Println(math.MaxUint64)
	// var x int32 = 1 << 40
	// fmt.Println(x)

	const (
		X = iota * 10
		Y
		_
		Z
	)
	fmt.Println(X, Y, Z)
}
