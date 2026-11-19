package main

import "fmt"

func main() {
	m := [][]int{{1, 2}, {3, 4}}

	rasa := make([][]int, len(m))
	copy(rasa, m)
	rasa[0][0] = 99
	fmt.Println(m, rasa)

	n := [][]int{{1, 2}, {3, 4}}
	funda := make([][]int, len(n))
	for i := range n {
		funda[i] = make([]int, len(n[i]))
		copy(funda[i], n[i])
	}
	funda[0][0] = 99
	fmt.Println(n, funda)
}
