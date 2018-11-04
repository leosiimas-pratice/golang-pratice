package main

import(
  "fmt"
)

func main(){
  funcESalario := map[string]float64{
    "João Paulo": 156.25,
    "Maria Roberto": 126.76,
    "José Paulo": 256.99,
  }

  funcESalario["Jose"] = 1234.44

  for nome, salario := range funcESalario{
    fmt.Println(nome, salario)
  }

  fmt.Println(funcESalario)
}
