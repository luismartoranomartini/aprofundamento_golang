package main

import "fmt"

func main() {
	s := "ir para casa    "
	fim := -1
	resultado := ""
	encontrouEspaco := false

	fmt.Printf("String original: %q (tamanho: %d)\n\n", s, len(s))

	for i := len(s) - 1; i >= 0; i-- {
		fmt.Printf("Volta i=%d | caractere=%q | fim antes=%d\n", i, s[i], fim)

		if s[i] != ' ' && fim == -1 {
			fim = i + 1
			fmt.Printf("    -> [1º if] Achou o fim da palavra! fim agora vale %d\n", fim)
		}
		if s[i] == ' ' && fim != -1 {
			resultado = s[i+1 : fim]
			encontrouEspaco = true
			fmt.Printf("   -> [2º if] Achou espaço antes da palavra! Fatiando s[%d:%d] = %q", i+1, fim, resultado)
			fmt.Println("  -> Dando break (não precisa olhar o resto à esquerda)")
			break
		}
		if !encontrouEspaco {
			if fim != -1 {
				resultado = s[:fim]
				fmt.Printf("[Fora do laço] Não achou espaço antes, mas fim=%d. Fatiando s[:%d] = %q\n", fim, fim, resultado)
			} else {
				resultado = ""
				fmt.Printf("[Fora do laço] fim continua -1. Resultado = %q\n", resultado)
			}
		}
	}
}
