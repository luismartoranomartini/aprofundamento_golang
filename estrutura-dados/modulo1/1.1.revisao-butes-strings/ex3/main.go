package main

import "fmt"

func main() {
	// n1, n2, n3, n4 := "Zebra", "abelha", "10", "9"

	palavras := []string{"Zebra", "abelha", "10", "9"}

	fmt.Println("Zebra" < "abelha", "10" < "9")

	for _, palavra := range palavras {
		bytes := []byte(palavra)
		fmt.Printf("Palavra: %-6s | Slice []byte: %v\n", palavra, bytes)
		for i := 0; i < len(palavra); i++ {
			b := palavra[i]
			fmt.Printf("byte[%d] = %3d (tipo: %T)\n", i, b, b)
		}
	}

}
