package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomePessoa() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNomePessoa(nomeCompleto string) {
	nomePessoa := strings.Split(nomeCompleto, " ")
	p.nome = nomePessoa[0]
	p.sobrenome = nomePessoa[1]
}

func main() {
	p1 := pessoa{"João", "Cesar"}

	fmt.Println(p1.getNomePessoa())

	p1.setNomePessoa("Maria Luiza")
	fmt.Println(p1.getNomePessoa())
}
