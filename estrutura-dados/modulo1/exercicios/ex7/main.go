package main

import "fmt"

func main() {
	palavra := "abcdef"
	palavraInvertida := []byte(palavra)

	for i := 0; i < len(palavra)/2; i++ {
		j := len(palavra) - 1 - i
		palavraInvertida[i], palavraInvertida[j] = palavraInvertida[j], palavraInvertida[i]
	}
	palavra = string(palavraInvertida)
	fmt.Println(palavra)

}
