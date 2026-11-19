package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a
	b[0] = 99
	fmt.Println(a, b)
	fmt.Println(&a[0] == &b[0])

	c := make([]int, 0, 3)
	n := copy(c, a)
	c[1] = 77
	fmt.Println(n, a, c)
	fmt.Println(&a[0] == &c[0])

	fmt.Println(min(len(c), len(a)))

	d := make([]int, 2)
	n = copy(d, a)
	fmt.Println(n, d)

	var e []int
	n = copy(e, a)
	fmt.Println(n, e)

	x := [3]int{1, 2, 3}
	y := x
	y[0] = 99
	fmt.Println(x, y)
}
