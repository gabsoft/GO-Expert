// Type assertion
package main

// func Soma(m map[string]int) int {
// 	var soma int
// 	for _, v := range m {
// 		soma += v
// 	}
// 	return soma
// }

// func SomaFloat(m map[string]float64) float64 {
// 	var soma float64
// 	for _, v := range m {
// 		soma += v
// 	}
// 	return soma
// }

type MyNumber int

// Constraint, ~ significa que aceita o tipo MyNumber que é um int, sem o ~ nao funcionaria
type Number interface {
	~int | ~float64
}

// Generics
func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

// GO Constraints

// func Compara[T comparable](a T, b T) bool {
// 	if a == b {
// 		return true
// 	}
// 	return false
// }

func Compara[T Number](a T, b T) bool {
	return a == b
}

func main() {
	m := map[string]int{"Gabriel": 1000, "Almeida": 1000}
	m2 := map[string]float64{"Gabriel": 1000.20, "Almeida": 1000.30}

	m3 := map[string]MyNumber{"Gabriel": 1000, "Almeida": 1000}

	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))

	println(Compara(10, 10.5))

}
