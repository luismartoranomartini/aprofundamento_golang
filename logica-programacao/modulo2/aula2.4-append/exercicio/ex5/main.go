package main

import "fmt"

func somas(s []int, k int) []int {
	if k <= 0 || k > len(s) {
		return nil
	}

	var res []int

	somaAtual := 0
	for i := 0; i < k; i++ {
		somaAtual += s[i]
	}
	res = append(res, somaAtual)
	for i := k; i < len(s); i++ {
		somaAtual += s[i] - s[i-k]
		res = append(res, somaAtual)
	}
	return res

}

func main() {
	seq := []int{2, 1, 5, 1, 3, 2}

	fmt.Println(somas(seq, 3))
	fmt.Println(somas(seq, 6))
	fmt.Println(somas(seq, 0))
	fmt.Println(somas(seq, 7))
	fmt.Println(somas([]int{-3, -1, -2}, 2))
}
