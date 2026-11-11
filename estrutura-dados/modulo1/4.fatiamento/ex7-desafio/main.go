package main

import "fmt"

func conta(s, sub string) int {

	if sub == "" {
		return -1
	}

	count := 0
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			count++
			i += len(sub) - 1
		}
	}
	return count
}

func main() {
	s := "banana"
	fmt.Println(conta(s, "an"))      // 2
	fmt.Println(conta(s, "x"))       // 0
	fmt.Println(conta(s, "na"))      // 2
	fmt.Println(conta("aaaa", "aa")) // 2
	fmt.Println(conta(s, ""))

	fmt.Println(len(s))
}
