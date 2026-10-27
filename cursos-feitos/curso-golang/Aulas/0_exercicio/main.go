package main

import (
    "fmt"
    "time"
    "exercicio1/model" // Ajuste para o nome do seu módulo/pasta
)

func main() {
    // Chamando a sua função que agora está perfeita
    compra, err := model.NewCompra("Mercado Central", time.Now(), []string{"Arroz", "Feijão", "carne"})

    if err != nil {
        fmt.Println("Erro ao criar compra:", err)
        return
    }

    fmt.Println("Compra criada com sucesso!")
    fmt.Println("Mercado:", compra.Mercado)
    fmt.Println("Quantidade de itens:", len(compra.Itens))
	fmt.Println(compra)
}
