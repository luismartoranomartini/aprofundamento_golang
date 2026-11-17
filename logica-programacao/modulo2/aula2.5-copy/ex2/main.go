package main

import "fmt"

func main() {
	src := []int{1, 2, 3}

	var a []int
	n := copy(a, src)
	fmt.Println("dst nil: n=", n, "dst =", a)

	b := make([]int, 0, 5)
	n = copy(b, src)
	fmt.Println("len 0 e cap 5: n=", n, "dst =", b, "cap=", cap(b))

	c := make([]int, len(src))
	n = copy(c, src)
	fmt.Println("len = len(src): n=", n, "dst =", c)
}
