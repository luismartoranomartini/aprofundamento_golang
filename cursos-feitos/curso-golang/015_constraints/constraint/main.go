// Package main
package main

import "fmt"

type myInt int

type Number interface {
	~int | float64
}

func SomaNumeros[T Number](a, b T) T {
	return a + b

}

func main() {
	var (
		x1 = int(100)
		x2 = int(200)

		y1 = float64(100.5)
		y2 = float64(200.5)

		z1 = myInt(100)
		z2 = myInt(200)
	)

	fmt.Println(SomaNumeros(x1, x2))
	fmt.Println(SomaNumeros(y1, y2))
	fmt.Println(SomaNumeros(float64(x1), y2))
	fmt.Println(SomaNumeros(z1, z2))

}
