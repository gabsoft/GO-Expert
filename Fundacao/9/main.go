package main

import "fmt"

func main() {
	fmt.Println(sum(1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1))
}

func sum(numeros ...int) int {
	resultado := 0
	for _, numero := range numeros {
		resultado += numero
	}
	return resultado
}
