package main

import (
	"fmt"
)

func main() {
	totalSegundos := 3665
	horas := totalSegundos / 3600
	restoSegundos := totalSegundos % 3600
	minutos := restoSegundos / 60
	segundos := restoSegundos % 60

	fmt.Printf("%02d:%02d:%02d\n", horas, minutos, segundos)
}
