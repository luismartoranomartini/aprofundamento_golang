package main

import "fmt"

func main() {
	// fallthrough: executa o corpo do case seguinte SEM testá-lo
	switch 1 {
	case 1:
		fmt.Println("um")
		fallthrough
	case 2:
		fmt.Println("dois (sem testar)")
	case 3:
		fmt.Println("três")
	}

	// default no meio: só vale se nenhum outro casar
	switch 5 {
	default:
		fmt.Println("default")
	case 5:
		fmt.Println("um")
		fallthrough
	}
}
