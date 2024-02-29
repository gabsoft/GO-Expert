package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}

	// tamanho, err := f.WriteString("Olá, mundo!")
	tamanho, err := f.Write([]byte("Olá, mundo! Olá, mundo! Olá, mundo! Olá, mundo! Olá, mundo! Olá, mundo!"))

	if err != nil {
		panic(err)
	}

	fmt.Printf("Tamanho: %d bytes\n", tamanho)
	f.Close()

	// abrir arquivo
	// arquivo, err := os.Open("test.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// leitura
	arquivos, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println((string(arquivos)))

	// leitura de pouco em pouco abrindo o arquivo

	arquivo, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}

	// Buffers

	reader := bufio.NewReader(arquivo)
	buffet := make([]byte, 10)
	for {
		n, err := reader.Read(buffet)
		if err != nil {
			break
		}
		fmt.Println(string(buffet[:n]))
	}

	err = os.Remove("test.txt")
	if err != nil {
		panic(err)
	}
}
