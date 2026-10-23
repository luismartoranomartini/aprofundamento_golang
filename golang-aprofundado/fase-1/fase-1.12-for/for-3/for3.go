package main

import "fmt"

func main() {
	// for i := 0; i < 2; i++ {
	//
	// }
	// fmt.Println(i)

	n := 2
	voltas := 0
	for range n {
		n = 5
		voltas++
	}
	fmt.Println(voltas, n)

	m := 2
	voltas = 0
	for i := 0; i < m; i++ {
		m = 4
		voltas++
	}
	fmt.Println(voltas, m)
}
