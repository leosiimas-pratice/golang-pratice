package main

import "fmt"

type nota float64

func (n nota) conceito(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func notaParaConceito(n nota) string {
	if n.conceito(9, 10) {
		return "A"
	} else if n.conceito(7, 8.9) {
		return "B"
	} else if n.conceito(5, 6.9) {
		return "C"
	} else if n.conceito(3, 4.9) {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.5))
	fmt.Println(notaParaConceito(5.5))
	fmt.Println(notaParaConceito(2.5))
}
