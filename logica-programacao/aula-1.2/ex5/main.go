package main

import "fmt"

func main() {
	anoBissexto := 2023

	if anoBissexto%4 == 0 {
		if anoBissexto%100 == 0 {
			if anoBissexto%400 == 0 {
				fmt.Println("É bissexto")
			} else {
				fmt.Println("Não é bissexto")
			}
		} else {
			fmt.Println("É bissexto")
		}

	} else {
		fmt.Println("não é bissexto")
	}
}
