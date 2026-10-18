package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	for _, v := range a {
		a[2] = 10
		fmt.Print(v, " ")
	}
	fmt.Println()

	s := []int{1, 2, 3}
	for _, v := range s {
		s[2] = 10
		fmt.Print(v, " ")
	}
	fmt.Println()
}
