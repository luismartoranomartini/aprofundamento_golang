package main

import "fmt"

func main() {
	var meuMap = map[string]any{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
	}
	fmt.Println(meuMap)

	for chave, valor := range meuMap {
		fmt.Println(chave, valor)
	}

	fmt.Println("--------------------")

	var meuMap2 = map[string]any{}
	meuMap2["chave"] = "valor"

	fmt.Println(meuMap2)

	fmt.Println("---------------------------")

	var meuMap3 = make(map[string]any) // nil
	meuMap3["chave"] = "valor"

	var meuvalor, ok = meuMap3["chave"].(int)
	if ok {
		fmt.Println(meuvalor)
	} else {
		fmt.Println("valor não encontrado")
	}
}
