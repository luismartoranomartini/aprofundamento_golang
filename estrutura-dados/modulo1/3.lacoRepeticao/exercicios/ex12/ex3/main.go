package main

import "fmt"

func main() {
	i := 0
	x := i++
	fmt.Println(x) //syntax error: unexpected ++ at end of statement
}

/* Previsão: Erro de compilação (syntax error: unexpected ++, expecting expression ou similar).

Por quê: Diferente de linguagens como C, C++ ou JavaScript, em Go o i++ é uma instrução (statement) e não uma expressão (expression). Ele não retorna um valor para ser guardado dentro de outra variável (x := ...); ele precisa ficar sozinho em uma linha. */
