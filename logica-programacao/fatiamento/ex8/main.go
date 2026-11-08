package main

import "fmt"

func firstUniqueChar(s string) int {
	m := [256]int{}

	for i := range s {
		cur := s[i]
		m[cur-'a']++
	}

	for i := range s {
		cur := s[i]
		if m[cur-'a'] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	palavra := "ddAee"
	fmt.Println(firstUniqueChar(palavra))
}
