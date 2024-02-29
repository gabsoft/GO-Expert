package main

func main() {
	// Condicionais

	a := 10
	b := 20

	if a > b {
		println("A é maior que B")
	} else {
		println("B é maior que A")
	}

	switch a {
	case 1:
		println("A é 1")
	case 2:
		println("A é 2")
	case 3:
		println("A é 3")
	default:
		println("A não é nenhum dos valores")
	}

}
