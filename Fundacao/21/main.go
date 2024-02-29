package main

import (
	"fmt"

	"github.com/gabsoft/golang-pacotes/matematica"
)

func main() {
	s := matematica.Soma(10, 20)

	fmt.Printf("Resultado: %v\n", s)
	fmt.Println(matematica.A)
}
