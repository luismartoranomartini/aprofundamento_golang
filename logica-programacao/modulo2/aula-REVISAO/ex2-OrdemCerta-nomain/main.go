package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40}
	i := 1
	v := 15

	fmt.Println("início:", s, len(s))

	s = append(s, 0)

	fmt.Println("append:", s, len(s))

	copy(s[i+1:], s[i:])
	fmt.Println("copy:", s, len(s))

	s[i] = v
	fmt.Println("gravar: ", s, len(s))
}
