package main

import "fmt"

func main() {
	// Tipos aceitos como chave: string, int, bool, rune, array de tipos aceitos...

	porLetra := map[rune]int{'a': 1, 'b': 2}
	porPar := map[[2]int]string{{0, 0}: "origem", {1, 2}: "ponto"}
	porBool := map[bool]string{true: "sim", false: "não"}

	fmt.Println(porLetra['a'], porPar[[2]int{1, 2}], porBool[3 > 2])

	// O valor pode ser de qualquer tipo, inclusive slice.
	grupos := map[string][]int{"pares": {2, 4}, "ímpares": {1, 3}}
	fmt.Println(grupos["pares"], len(grupos["ímpares"]), grupos["nada"] == nil)

}
