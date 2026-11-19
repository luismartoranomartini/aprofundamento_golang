package main

import "fmt"

func main() {
	linha := []int{0, 0}
	m := [][]int{linha, linha}
	m[0][0] = 5
	fmt.Println(m)
	fmt.Println(&m[0][0] == &m[1][0])

	n := make([][]int, 2)
	for i := range n {
		n[i] = make([]int, 2)
	}
	n[0][0] = 5
	fmt.Println(n)
	fmt.Println(&n[0][0] == &n[1][0])
}
