package main

import "fmt"

// Interfaces vazias aceitam qualquer tipo, tomar cuidado com o uso, é melhor evitar e usar generics
type x interface{}

func main() {
	var x interface{} = 10
	var y interface{} = "Hello World"

	showType(x)
	showType(y)

}

func showType(t interface{}) {
	fmt.Printf("O tipo da variavel é: %T e o valor é %v\n", t, t)
}
