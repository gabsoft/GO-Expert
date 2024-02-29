package main

func main() {
	// Memoria -> Endereco -> Valor
	// Ponteiro === endereco na memoria
	// Porque usar ponteiro? Muitas vezes voce pode nao querer trabalhar com o valor local e sim no endereco da memoria

	a := 10

	var ponteiro *int = &a
	*ponteiro = 20
	b := &a
	*b = 30
	println(b)
	println(*b)
	println(a)

}
