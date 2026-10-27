package main

import "fmt"

type MyInt int

type Number interface {
	~int | float64
}

func Soma[T Number](a, b T) T {
	return a + b
}

func main() {
	var (
		x1 = int(100)
		x2 = int(200)

		y1 = float64(100.5)
		y2 = float64(200.5)

		z1 = MyInt(100)
		z2 = MyInt(200)
	)

	fmt.Println("Soma de x1 + x2 =", Soma(x1, x2))
	fmt.Println("Soma de y1 + y2 =", Soma(y1, y2))

	fmt.Println("Soma de z1 + z2 =", Soma(z1, z2))

	// fmt.Println("Soma de x1 + y2 =", Soma(x1, y2))
	// fmt.Println("Soma de x1 + x2 com int64:", Soma(int64(x1), int64(x2)))
}
