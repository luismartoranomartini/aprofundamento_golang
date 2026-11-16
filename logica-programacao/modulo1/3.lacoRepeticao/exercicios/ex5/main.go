package main

import "fmt"

func main() {
	texto := "São Paulo"
	cont := 0
	cont1 := 0

	for range texto {
		cont++
	}
	fmt.Println(cont)

	for i := 0; i < len(texto); i++ {
		cont1++
	}
	fmt.Println(cont1)
}
