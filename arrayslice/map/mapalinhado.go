package main

import (
  "fmt"
)

func main(){
  funcSlarioPLetra := map[string]map[string]float64{
    "G":{
      "Gustavo Lima": 15.66,
      "Gyasd Aasda": 1223.4,
    },
    "H":{
      "Handrey Lua": 123.5,
    },
    "L":{
      "Leonardo Abreu": 21312.44,
      "Leo": 1233.223,
    },
  }

  fmt.Println(funcSlarioPLetra)

  for letra, funcionarios := range funcSlarioPLetra{
    fmt.Println(letra)
    for nome, salario := range funcionarios{
      fmt.Println(nome, salario)
    }

  }

}
