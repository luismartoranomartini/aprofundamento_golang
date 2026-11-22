package main

import "fmt"

func temDuplicado(s []int) bool {
	duplicatas := make(map[int]bool)
	existe := false

	for _, num := range s {
		if duplicatas[num] == true {
			existe = true
		}
		duplicatas[num] = true
	}

	return existe
}

func main() {
	fmt.Println(temDuplicado([]int{1, 2, 3}))
	fmt.Println(temDuplicado([]int{1, 2, 1}))
	fmt.Println(temDuplicado([]int{}))
	fmt.Println(temDuplicado(nil))
}
