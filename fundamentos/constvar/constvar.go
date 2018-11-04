package main

import (
  "fmt"
  m "math"
)

func main(){
  const PI float64 = 3.1415
  var raio = 3.2 // tipo(float64) inferido pelo compilador

  // forma reduzida de criar uma var
  area := PI * m.Pow(raio,2)
  fmt.Println("Area da circunferencia é ", area)

  const (
    a = 1
    b = 2
  )
  var(
    c = 3
    d = 4
  )
  fmt.Println(a,b,c,d)

  //  var  tipo   recebem, estanciando numa mesma linha varias variaveis
  var e, f bool = true, false
  fmt.Println(e,f)
  // formato reduzido do anterior
  g, h, i := 2, false, "epa!"
  fmt.Println(g,h,i)

}
