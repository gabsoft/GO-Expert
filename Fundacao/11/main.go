package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	gabriel := Cliente{
		Nome:  "Gabriel",
		Idade: 25,
		Ativo: true,
	}

	gabriel.Ativo = false

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", gabriel.Nome, gabriel.Idade, gabriel.Ativo)

	fmt.Println(gabriel)
}
