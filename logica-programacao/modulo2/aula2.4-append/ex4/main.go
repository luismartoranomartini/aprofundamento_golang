package main

import "fmt"

func main() {
	a := make([]int, 3, 4)
	a[0], a[1], a[2] = 1, 2, 3

	b := append(a, 100)
	c := append(a, 200)

	fmt.Println("b[3] =", b[3], "| c[3] =", c[3])
	fmt.Println("mesmo array (a, b)?", &a[0] == &b[0])

	d := []int{1, 2, 3}
	e := append(d, 4)
	e[0] = 99

	fmt.Println("d =", d)
	fmt.Println("e =", e)
	fmt.Println("mesmo array (d, e):", &d[0] == &e[0])
}
