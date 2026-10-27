package main

import "fmt"

func notaParaConceito(n float64) string {
	// var nota = int(n)
	var nota = float64(n)
	switch nota {
	case 10:
		fallthrough // para entrar nos cases
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println(notaParaConceito(10))

}
