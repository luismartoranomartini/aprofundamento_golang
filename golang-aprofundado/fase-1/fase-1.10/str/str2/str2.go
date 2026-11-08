package main

import "fmt"

func main() {
	s := "ação"
	fmt.Println(s[1:3] == "ç")
	fmt.Println(s[1:2])
	fmt.Printf("%q\n", s[1:2])
	fmt.Println(len("a"), len("ç"), len("世"), len("😀"))

	s1 := "olá"
	fmt.Println(len(s1))
	fmt.Println(len([]rune(s1)))

	s2 := "casa"
	r := []rune(s2)
	r[0] = 'm'
	fmt.Println(s2)
	fmt.Println(string(r))

}
