package main

import "fmt"

func main() {
	total := func() int {
		return sum(1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1) * 2
	}()

	fmt.Println(total)
}

func sum(numeros ...int) int {
	resultado := 0
	for _, numero := range numeros {
		resultado += numero
	}
	return resultado
}
