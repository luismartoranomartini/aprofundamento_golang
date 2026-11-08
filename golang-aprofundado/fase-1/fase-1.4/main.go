package main

import "fmt"

func main() {
	var s []int
	s = append(s, 1, 2, 3)
	fmt.Println(s, len(s), cap(s))

	a := []int{1, 2, 3, 4}
	b := a[:2]
	b = append(b, 99)
	fmt.Println(a, b)

	c := append(a, 5)
	c[0] = 1000
	fmt.Println(a, c)

	s1 := make([]int, 3, 5)
	t := append(s1, 1)
	u := append(s1, 2)
	fmt.Println(s1, t, u)
}
