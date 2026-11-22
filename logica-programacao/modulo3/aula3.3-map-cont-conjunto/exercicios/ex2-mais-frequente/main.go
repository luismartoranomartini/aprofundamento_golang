package main

import "fmt"

func maisFrequente(s []int) (int, bool) {
	contagem := make(map[int]int)

	for _, v := range s {
		contagem[v]++
	}

	var vantagem int
	maiorContagem := 0

	for num, v := range contagem {
		if v > maiorContagem {
			maiorContagem = v
			vantagem = num
		}
	}

	if maiorContagem > 0 {

		return vantagem, true
	}

	return 0, false
}

func main() {
	fmt.Println(maisFrequente([]int{5, 5, 7}))
	fmt.Println(maisFrequente([]int{4, 9, 9, 9, 4}))
	fmt.Println(maisFrequente([]int{}))
}
