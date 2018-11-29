package main

import (
	"fmt"
)

type impremivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x impremivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa impremivel = pessoa{"ROBERTO", "Silva"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = produto{"lapis", 6.90}
	fmt.Println(coisa.toString())
	imprimir(coisa)
}
