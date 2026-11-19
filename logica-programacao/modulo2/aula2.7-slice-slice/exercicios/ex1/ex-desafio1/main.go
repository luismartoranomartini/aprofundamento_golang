package main

import "fmt"

// somaLinhas devolve UM número por linha: a soma da linha.
func somaLinhas(m [][]int) []int {
	// acumula guarda o RESULTADO: uma soma por linha. Começa nil (vazio).
	var acumula []int
	// soma é o acumulador de UMA linha: vai somando os elementos da linha da vez.
	var soma int

	for i := 0; i < len(m); i++ { // laço de fora: uma volta por LINHA
		// Zera o acumulador a cada linha nova. Sem isso, a soma da linha 1
		// continuaria somando em cima da linha 0.
		soma = 0
		for j := 0; j < len(m[i]); j++ { // laço de dentro: uma volta por ELEMENTO da linha i
			soma += m[i][j] // aqui só se acumula; a soma ainda NÃO está pronta
		}
		// Fim do laço de dentro: a linha inteira foi somada. Só agora a soma está pronta.
		// O append fica no laço de fora porque o resultado tem UM item por linha.
		acumula = append(acumula, soma)
	}

	return acumula
}

// somasParciais devolve UM número por ELEMENTO: a soma até aquele elemento, linha a linha.
func somasParciais(m [][]int) []int {
	var res []int
	for i := 0; i < len(m); i++ {
		soma := 0
		for j := 0; j < len(m[i]); j++ {
			soma += m[i][j]
			// Aqui o append fica DENTRO: cada elemento produz um item no resultado.
			res = append(res, soma)
		}
	}
	return res
}

// achata devolve todos os elementos da matriz numa lista só, em ordem.
func achata(m [][]int) []int {
	var res []int
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			// De novo DENTRO: cada elemento vira um item do resultado.
			res = append(res, m[i][j])
		}
	}
	return res
}

func main() {
	a := [][]int{{2, 2}, {5}}
	fmt.Println(somaLinhas(a))    // [6 9 6]       3 linhas -> 3 itens
	fmt.Println(somasParciais(a)) // [1 3 6 4 9 6] 6 elementos -> 6 itens
	fmt.Println(achata(a))        // [1 2 3 4 5 6] 6 elementos -> 6 itens

	b := [][]int{{}, {7, 8}}
	fmt.Println(somaLinhas(b)) // [0 15]

	fmt.Println(somaLinhas(nil)) // []
}
