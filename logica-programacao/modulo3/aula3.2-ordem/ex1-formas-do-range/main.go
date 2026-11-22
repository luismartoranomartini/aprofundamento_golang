package main

import "fmt"

func main() {
	notas := map[string]int{"Ana": 8, "Bia": 10, "Caio": 6}

	soma := 0
	letras := 0
	for k, v := range notas {
		soma += v
		letras += len(k)
	}
	fmt.Println("Soma:", soma, "letras:", letras)

	n := 0
	for k := range notas {
		if len(k) == 3 {
			n++
		}
	}
	fmt.Println("chaves com 3 letras", n)

	maior := 0
	for _, v := range notas {
		if v > maior {
			maior = v
		}
	}
	fmt.Println("maior nota:", maior)

	menor := 10
	for _, v := range notas {
		if v < menor {
			menor = v
		}
	}
	fmt.Println("menor nota:", menor)

	mediana := 0
	for _, v := range notas {
		if v > menor && v < maior {
			mediana = v
		}
	}
	fmt.Println("mediana", mediana)

	for _, v := range notas {
		if v > menor && v < maior {
			fmt.Println("média:", v)
		}
	}

	fmt.Println(notas)

	for k, v := range notas {
		fmt.Println(k, v)
	}

}
