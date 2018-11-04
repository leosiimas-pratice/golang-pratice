package main

import (
  "fmt"
)

func notaParaConceito(n float64) string {
  var nota = n

  switch true {
  case nota <= 10 && nota >= 9:
    return "A"
  case nota < 9 && nota >= 8:
    return "B"
  case nota < 8 && nota >= 5:
    return "C"
  case nota < 5 && nota >= 3:
    return "D"
  case nota < 3 && nota >= 0:
    return "D"
  default:
    return "Nota Invalida"
  }
}

func main(){
  fmt.Println(notaParaConceito(50))
  fmt.Println(notaParaConceito(10))
  fmt.Println(notaParaConceito(9.5))
  fmt.Println(notaParaConceito(6))

}
