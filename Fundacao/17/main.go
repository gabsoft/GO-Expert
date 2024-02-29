package main

type Conta struct {
	saldo int
}

func (c *Conta) simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}

// Especie de construtor, qual a motivacao de usar isso?
func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func main() {
	conta := Conta{saldo: 1000}
	conta.simular(200)
	println(conta.saldo)
}
