package main

import "fmt"

func agruparPorResto(num []int) map[int][]int {
	grupo := make(map[int][]int)

	for _, n := range num {
		resto := n % 3
		grupo[resto] = append(grupo[resto], n)
	}
	return grupo
}

func main() {
	entrada := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(agruparPorResto(entrada))
}
