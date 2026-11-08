package dobro

import "fmt"

func Dobro(n int) int { return n * 2 }

func main() {
	numero := 2
	Dobro(numero)
	fmt.Println(numero)
}
