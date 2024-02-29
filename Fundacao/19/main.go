// Type assertion
package main

import "fmt"

// Interfaces vazias aceitam qualquer tipo, tomar cuidado com o uso, é melhor evitar e usar generics
type x interface{}

func main() {
	var minhaVar interface{} = "Gabriel Almeida"

	fmt.Println(minhaVar.(string))
	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res é %v e o valor de ok é %v\n", res, ok)
	res2 := minhaVar.(int)
	fmt.Printf("O valor de res2 é %v\n", res2)

}
