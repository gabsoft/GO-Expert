package matematica

// Sempre para exportar sao os metodos comecando com maiusculo

func Soma[T int | float64](a, b T) T {
	return a + b
}

var A int = 10
