package main

import "fmt"

func main() {
	// Lista inicial com 5 produtos
	produtos := []string{"Arroz", "Feijão", "Macarrão", "Azeite", "Sal"}
	fmt.Printf("1. Slice original:%v\n\n", produtos)

	// Queremos remover "Macarrão", que está no índice 2
	i := 2
	fmt.Printf("Elemento a remover:'%s' (índice i = %d)\n\n", produtos[i], i)

	// Fatias separadas
	antes := produtos[:i]    // do início até i-1
	depois := produtos[i+1:] // de i+1 até o fim

	fmt.Printf("2. Parte antes  [:i]: %v (índices 0 a %d)\n", antes, i-1)
	fmt.Printf("3. Parte depois [i+1:]: %v (índices %d ao fim)\n\n", depois, i+1)

	// Juntando as duas metades sem incluir produtos[i]
	produtos = append(antes, depois...)

	fmt.Printf("4. Slice final: %v\n", produtos)
}
