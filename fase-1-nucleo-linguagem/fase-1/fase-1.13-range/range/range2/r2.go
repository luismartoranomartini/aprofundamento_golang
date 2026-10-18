package main

import "fmt"

func main() {
	// array: range copia o array
	a := [3]int{1, 2, 3}
	for i, v := range a {
		if i == 0 {
			a[2] = 100
		}
		fmt.Print(v, " ")
	}
	fmt.Println(a)

	// ponteiro para array: não copia
	b := [3]int{1, 2, 3}
	for i, v := range &b {
		if i == 0 {
			b[2] = 100
		}
		fmt.Print(v, " ")
	}
	fmt.Println(b)

	// slice: elementos lidos na hora; len fixado no início
	s := []int{1, 2, 3}
	for i, v := range s {
		if i == 0 {
			s[2] = 100
			s = append(s, 7)
		}
		fmt.Print(v, " ")
	}
	fmt.Println(s)
}
