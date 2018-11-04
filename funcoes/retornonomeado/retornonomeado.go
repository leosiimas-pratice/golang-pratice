package main

import "fmt"

func retornonomeado(p1,p2 int)(segundo, primeiro int){
  segundo = p2
  primeiro = p1
  return
}

func main(){

  x,y := retornonomeado(5, 7)
  fmt.Println(x,y)
}
