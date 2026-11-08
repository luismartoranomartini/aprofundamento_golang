package main

import (
	"errors"
	"fmt"
)

func MaiorEIndice(numeros []int) (maior, indice int, err error) {
	if len(numeros) == 0 {
		return 0, -1, errors.New("o slice não pode estar vazio")
	}
	maior = numeros[0]
	indice = 0

	for i := 1; i < len(numeros); i++ {
		if numeros[i] > maior {
			maior = numeros[i]
			indice = i
		}
	}
	return maior, indice, nil
}

func main() {
	valores := []int{7 - 5, -2, -8}

	maior, pos, err := MaiorEIndice(valores)
	if err != nil {
		fmt.Println("erro:", err)
		return
	}
	fmt.Printf("Maior valor: %d\n", maior)
	fmt.Printf("Primeira posição (índice): %d\n", pos)
}
