package main

import "fmt"

func main() {
	s := "a1b2c3"

	soma := 0

	for i := 0; i < len(s); i++ {
		num := s[i]
		if num >= '1' && num <= '9' {
			soma += int(num - '0')
		}
	}
	fmt.Println(soma)
}
