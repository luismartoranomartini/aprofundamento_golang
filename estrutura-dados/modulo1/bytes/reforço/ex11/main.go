package main

import "fmt"

func invete(s string) string {
	b := []byte(s)
	for i := 0; i < len(s)/2; i++ {
		j := len(b) - 1 - i
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func main() {
	// palavra := "12345"
	// palavraInvertida := []byte(palavra)
	//
	// for i := 0; i < len(palavra)/2; i++ {
	// 	j := len(palavra) - 1 - i
	// 	palavraInvertida[i], palavraInvertida[j] = palavraInvertida[j], palavraInvertida[i]
	// }
	//
	// palavra = string(palavraInvertida)
	// fmt.Println(palavra)
	s := "12345"
	v := ""
	x := "x"
	fmt.Println(invete(s))
	fmt.Printf("%q", invete(v))
	fmt.Println(invete(x))

}
