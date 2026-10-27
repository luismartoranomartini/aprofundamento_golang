package model

type Auto struct {
	Ano    int
	Placa  string
	Modelo string
}

type Moto struct {
	Auto
	Cilindradas int
}

type Carro struct {
	Auto
	Portas int
	Power  int
	Ar     bool
}
