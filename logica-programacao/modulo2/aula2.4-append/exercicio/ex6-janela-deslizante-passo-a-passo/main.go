package main

import "fmt"

func main() {
	s := []int{2, 1, 5, 1, 3, 2}
	k := 3

	fmt.Printf("Entrada: s = %v | Tamanho da janela (k) = %d\n\n", s, k)

	if k <= 0 || k > len(s) {
		fmt.Println("k inválido, resultado: nil")
		return
	}

	var res []int

	// 1. Primeira janela (índices 0 até k-1)
	somaAtual := 0
	for i := 0; i < k; i++ {
		somaAtual += s[i]
	}
	res = append(res, somaAtual)
	fmt.Printf("[Inicial] Janela %v -> somaAtual: %d | res: %v (len=%d, cap=%d)\n\n",
		s[:k], somaAtual, res, len(res), cap(res))

	// 2. Deslizamento da janela
	for i := k; i < len(s); i++ {
		sai := s[i-k]
		entra := s[i]
		somaAnterior := somaAtual

		somaAtual += entra - sai
		res = append(res, somaAtual)

		fmt.Printf("[i=%d] Sai s[%d]=%d | Entra s[%d]=%d\n", i, i-k, sai, i, entra)
		fmt.Printf("      Conta: %d + %d - %d = %d\n", somaAnterior, entra, sai, somaAtual)
		fmt.Printf("      Janela atual: %v | res agora: %v (len=%d, cap=%d)\n\n",
			s[i-k+1:i+1], res, len(res), cap(res))
	}

	fmt.Println("Resultado final:", res)
}
