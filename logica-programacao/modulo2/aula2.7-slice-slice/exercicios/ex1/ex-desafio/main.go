package main

func somaLinhas(m [][]int) []int {
	var acumula []int
	var soma int

	for i := 0; i < len(m); i++ {
		soma = 0
		for j := 0; j < len(m[i]); j++ {
			acumula = append(acumula, soma)
			soma += m[i][j]

			// acumula = append(acumula, soma)
			// soma += m[i][j]
		}
	}

	return acumula
}

func main() {

	// var acumula []int
	//
	// a := [][]int{{}, {7, 8}}
	//
	// for i := 0; i < len(a); i++ {
	// 	var soma int
	// 	for j := 0; j < len(a[i]); j++ {
	// 		acumula = append(acumula, soma)
	// 		soma += a[i][j]
	// 	}
	// 	fmt.Printf("[%d]", soma)
	// }
}
