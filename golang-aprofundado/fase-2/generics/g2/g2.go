package main

import "fmt"

func Map[T, U any](s []T, f func(T) U) []U {
	r := make([]U, 0, len(s))
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}

func Filter[T any](s []T, ok func(T) bool) []T {
	var r []T
	for _, v := range s {
		if ok(v) {
			r = append(r, v)
		}
	}
	return r
}

func Primeiro[T any](s []T) (T, bool) {
	var zero T
	if len(s) == 0 {
		return zero, false
	}
	return s[0], true
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(Map(nums, func(n int) int { return n * 2 }))
	fmt.Println(Map(nums, func(n int) string { return fmt.Sprint(n, "!") }))
	fmt.Println(Filter(nums, func(n int) bool { return n%2 == 0 }))
	v, ok := Primeiro([]string{})
	fmt.Printf("%q %v\n", v, ok)
	n, ok := Primeiro(nums)
	fmt.Println(n, ok)
}
