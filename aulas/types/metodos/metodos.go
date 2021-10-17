package main

import "fmt"

type humano interface {
	falar()
}

type pessoa struct {
	nome  string
	idade int
}

func (p *pessoa) falar() {
	fmt.Println("Meu nome é", p.nome)
}

func dizerAlgo(h humano) {
	h.falar()
}

func main() {
	p1 := pessoa{"Jhon Doe", 22}

	// Maneiras de executar a função
	p1.falar()
	(&p1).falar()
	dizerAlgo(&p1)
}
