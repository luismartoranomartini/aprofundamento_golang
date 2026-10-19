package main

import "fmt"

func mexe(n int, a [2]int, s []int, m map[string]int) {
	n = 100
	a[0] = 100
	s[0] = 100
	m["k"] = 100
}

func cresce(s []int) {
	s = append(s, 7)
}

func main() {

	n := 1
	a := [2]int{1, 2}
	s := []int{1, 2}
	m := map[string]int{"k": 1}
	mexe(n, a, s, m)
	fmt.Println(n, a, s, m)

	t := []int{1, 2}
	cresce(t)
	fmt.Println(t, len(t))
}
