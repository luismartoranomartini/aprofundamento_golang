package main

import "fmt"

func main() {
	palavraCrua := "GOLANG 2026"
	palavra := []byte(palavraCrua)

	for i := 0; i < len(palavra); i++ {
		if palavra[i] >= 'A' && palavra[i] <= 'Z' {
			palavra[i] += 'a' - 'A'
		}
	}
	palavraCrua = string(palavra)
	fmt.Println(palavraCrua)
}
