package main

import(
  "fmt"
)

func main(){
  i := 1

  // GO não tem aritmética de ponteiros
  p++
  
  var p *int = nil

  p = &i // pegando o endereço da variável
  fmt.Println(p,*p, i)

  *p++
  fmt.Println(i, &i)
}
