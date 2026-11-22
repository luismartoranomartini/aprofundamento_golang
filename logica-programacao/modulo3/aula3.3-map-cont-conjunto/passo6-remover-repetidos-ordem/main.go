package main

import "fmt"

func unicos(s []int) []int {
	visto := make(map[int]bool)
	var res []int

	for _, v := range s {
		jaVisto := visto[v]

		if !jaVisto {
			visto[v] = true
			res = append(res, v)
		}

		fmt.Println("v =", v, "| já visto antes?", jaVisto, "| res = res")
	}
	return res
}

func main() {
	fmt.Println("resposta", unicos([]int{3, 1, 3, 2, 1, 2}))
}
