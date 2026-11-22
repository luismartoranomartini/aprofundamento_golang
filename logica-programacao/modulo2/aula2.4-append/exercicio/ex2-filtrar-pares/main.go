package main

import "fmt"

func pares(s []int) []int {

	var par []int
	for _, p := range s {
		if p%2 == 0 {
			par = append(par, p)
		}
	}
	return par
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	b := []int{1, 3, 5}
	c := []int{}

	fmt.Println(pares(a))
	fmt.Println(pares(b))
	fmt.Println(pares(c))
	fmt.Println(pares(nil))
}
