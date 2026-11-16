package main

import "fmt"

func main() {
	a := []int{1, 2}

	a = append(a, 3)
	fmt.Println("um valor:           ", a)

	a = append(a, 4, 5, 6)
	fmt.Println("vários valores:     ", a)

	b := []int{7, 8, 9}
	a = append(a, b...)
	fmt.Println("com b...:           ", a)

	c := append(a)
	fmt.Println("sem valores:         ", len(c) == len(a), cap(c) == cap(a))

	var n []int
	n = append(n, 1, 2)
	fmt.Println("de nil:              ", n, n == nil)

	fmt.Printf("tipo:                 %T\n", n)

	/*
		s := []int{1, 2, 3}
		append(s, 4)
	*/

	/*
		a = append(a, b)
	*/
}
