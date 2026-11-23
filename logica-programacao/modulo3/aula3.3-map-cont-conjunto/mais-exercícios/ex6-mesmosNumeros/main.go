package main

import "fmt"

func mesmosNumeros(a, b []int) bool {
	contador := make(map[int]int)

	for _, x := range a {
		contador[x]++
	}

	for _, x := range b {
		contador[x]--
	}

	for _, v := range contador {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(mesmosNumeros([]int{1, 2, 2, 3}, []int{2, 3, 1, 2}))
	fmt.Println(mesmosNumeros([]int{1, 2, 2}, []int{1, 1, 2}))
	fmt.Println(mesmosNumeros([]int{}, []int{}))
}
