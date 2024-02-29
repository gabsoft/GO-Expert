package main

import "fmt"

func main() {
	salarios := map[string]int{"Wesley": 1999, "Joao": 2000, "Maria": 3000}

	fmt.Println(salarios["Wesley"])
	delete(salarios, "Wesley")
	salarios["Wes"] = 5000
	fmt.Println(salarios)

	// sal := make(map[string]int)
	sal1 := map[string]int{}
	sal1["Gabriel"] = 1000

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	}

	// Blank identifier
	for _, salario := range salarios {
		fmt.Printf("O salario é %d\n", salario)
	}

}
