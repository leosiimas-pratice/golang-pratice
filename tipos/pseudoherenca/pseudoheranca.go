package main

import (
	"fmt"
)

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro
	turboLigado bool
}

func main() {

	f := ferrari{}
	f.nome = "F40"
	f.turboLigado = true
	f.velocidadeAtual = 100

	fmt.Printf("O carro %s esta com o tubo ligado? %v \n", f.nome, f.turboLigado)

}
