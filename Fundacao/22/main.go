package main

func main() {
	// No go temos apenas o for

	for i := 0; i < 3; i++ {
		println(i)
	}

	numeros := []string{"zero", "um", "dois"}
	for k, v := range numeros {
		println(k, v)
	}

	i := 0
	for i < 3 {
		println(i)
		i++
	}

	for {
		println("loop infinito")
	}
}
