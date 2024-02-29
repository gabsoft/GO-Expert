package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

// Interface permite passar apenas métodos
type Pessoa interface {
	Desativar()
}

type Cliente struct {
	Nome    string
	Idade   int
	Ativo   bool
	Address Endereco
	Endereco
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {
	fmt.Printf("Desativando empresa %s\n", e.Nome)
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("Desativando cliente %s\n", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	gabriel := Cliente{
		Nome:  "Gabriel",
		Idade: 25,
		Ativo: true,
	}

	minhaEmpresa := Empresa{}

	Desativacao(gabriel)
	Desativacao(minhaEmpresa)

}
