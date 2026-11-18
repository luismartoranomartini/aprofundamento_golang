package main

import "fmt"

func unicosContando(s []int) (res []int, comparacoes int) {
	for _, v := range s {
		achou := false
		for _, x := range res {
			comparacoes++
			if x == v {
				achou = true
				break
			}
		}
		if !achou {
			res = append(res, v)
		}
	}
	return res, comparacoes
}

func main() {
	for _, n := range []int{10, 100, 1000} {
		s := make([]int, n)
		for i := range s {
			s[i] = i
		}
		_, c := unicosContando(s)
		fmt.Printf("n = %4d comparacoes = %6d\n", n, c)
	}
}
