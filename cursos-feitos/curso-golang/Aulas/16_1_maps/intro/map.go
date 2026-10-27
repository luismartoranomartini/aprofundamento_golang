package main

import (
	"fmt"
	"slices"
)

func main() {
	var meuMap = map[string]any{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
	}

	// fmt.Println(meuMap)

	chaves := make([]string, 0, len(meuMap))

	for chave := range meuMap {
		chaves = append(chaves, chave)
	}

	// sort.Strings(chaves)
	slices.Sort(chaves)

	for _, chave := range chaves {
		v := meuMap[chave]
		fmt.Println(chave, v)
	}
	fmt.Println("-----------------------")

	var meuMap2 = map[string]any{}
	meuMap2["chave"] = "valor"
	fmt.Println(meuMap2)
	// fmt.Println(meuMap2["chaave"])

	var meuValor, ok = meuMap2["chave"].(string)
	if ok {
		fmt.Println(meuValor)
	} else {
		fmt.Println("valor não encontrado")
	}
	// fmt.Println(meuValor)

	// casting
}
