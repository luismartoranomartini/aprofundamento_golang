package main

import "fmt"

func menosFrequente(s []int) (int, bool) {
	if len(s) <= 0 {
		return 0, false
	}

	contagem := make(map[int]int)

	for _, v := range s {
		contagem[v]++
	}

	var menos int
	menorContagem := len(s) + 1

	for num, v := range contagem {
		if menorContagem > v {
			menorContagem = v
			menos = num
		}
	}

	return menos, true
}

func main() {
	fmt.Println(menosFrequente([]int{5, 5, 7}))
	fmt.Println(menosFrequente([]int{4, 9, 9, 9, 4, 1}))
	fmt.Println(menosFrequente([]int{}))
	fmt.Println(menosFrequente(nil))
}
