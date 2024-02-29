package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome    string
	Idade   int
	Ativo   bool
	Address Endereco
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("Desativando cliente %s\n", c.Nome)
}

func main() {
	gabriel := Cliente{
		Nome:  "Gabriel",
		Idade: 25,
		Ativo: true,
	}

	gabriel.Ativo = false
	gabriel.Address.Cidade = "São Paulo"
	gabriel.Cidade = "São Paulo"

	gabriel.Desativar()
}
