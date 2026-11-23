package main

import "fmt"

func separarParImpar(entrada []int, par map[string][]int) map[string][]int {
	for _, n := range entrada {
		if n%2 == 0 {
			par["par"] = append(par["par"], n)
		} else {
			par["impar"] = append(par["impar"], n)
		}
	}
	return par
}

func main() {
	entrada := []int{1, 2, 3, 4, 5, 6, 7}
	par := make(map[string][]int)

	fmt.Println(separarParImpar(entrada, par))

	// for _, n := range entrada {
	// 	if n%2 == 0 {
	// 		par["par"] = append(par["par"], n)
	// 	} else {
	// 		par["impar"] = append(par["impar"], n)
	// 	}
	// }
	// fmt.Println(entrada, par)
}
