package main

import "fmt"

func contaMinusculas(s string) int {
	contMin := 0

	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			contMin++
		}
	}

	return contMin

}

func main() {
	fmt.Println(contaMinusculas("Go Lang 1"))
	fmt.Println(contaMinusculas("ABC"))
	fmt.Println(contaMinusculas("abc"))
	fmt.Println(contaMinusculas(""))
}
