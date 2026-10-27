package main

import "fmt"

func main() {

	imprimeMensagem("Olá, Luís")
	imprimeMensagem("Olá, Jonice")
	
}

func imprimeMensagem(msg string) {
	msg += ", bom dia! "
	fmt.Println(msg)
}

