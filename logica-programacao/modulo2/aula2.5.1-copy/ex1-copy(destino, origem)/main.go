package main

import "fmt"

func main() {
	origem := []int{1, 2, 3, 4, 5}

	// Caso 1: destino MENOR que a origem. Copia só len(destino) elementos.
	destino := make([]int, 3)
	n := copy(destino, origem)
	fmt.Println(n, destino) // 3 [1 2 3]

	// Caso 2: destino MAIOR que a origem. Copia len(origem) elementos; o resto não muda.
	curta := []int{9, 8}
	grande := []int{0, 0, 0, 0, 0}
	n = copy(grande, curta)
	fmt.Println(n, grande) // 2 [9 8 0 0 0]

	// Caso 3: destino nil tem len 0. Nada é copiado e não há erro.
	var vazio []int
	n = copy(vazio, origem)
	fmt.Println(n, vazio) // 0 []

	// Caso 4: copy olha o len do destino, não o cap. Aqui len 0 e cap 5: copia 0.
	reserva := make([]int, 0, 5)
	n = copy(reserva, origem)
	fmt.Println(n, reserva, len(reserva), cap(reserva)) // 0 [] 0 5

	// Caso 5: caso especial da linguagem: destino []byte e origem string.
	// (byte é um número inteiro de 0 a 255, visto no Módulo 1.)
	bs := make([]byte, 3)
	n = copy(bs, "golang")
	fmt.Println(n, bs, string(bs)) // 3 [103 111 108] gol

}
