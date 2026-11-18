package main

import "fmt"

func pares(s []int) []int {
	var res []int
	for _, v := range s {
		if v%2 == 0 {
			res = append(res, v)
		}
	}
	return res
}

func paresNoLugar(s []int) []int {
	res := s[:0]
	for _, v := range s {
		if v%2 == 0 {
			res = append(res, v)
		}
	}
	return res
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("novo:", pares(a), "|a depois:", a, "len=", len(a), "cap=", cap(a))

	b := []int{1, 2, 3, 4, 5, 6}
	r := paresNoLugar(b)
	fmt.Println("no lugar: ", r, "len r=", len(r), "|b depois:", b, "len=", len(b), "cap=", cap(b))
}
