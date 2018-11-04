package main

import(
  "fmt"
)

func main(){

  aprovados := make(map[int]string)

  aprovados[12345678978] = "Maria"
  aprovados[12345678178] = "Pedro"
  aprovados[12345678944] = "Ana"
  aprovados[12345232978] = "João"

  for cpf, nome := range aprovados{
    fmt.Printf("%s (CPF: %d)\n", nome, cpf)
  }

  fmt.Println(aprovados)
  delete(aprovados, 12345232978)
  fmt.Println(aprovados)

}
