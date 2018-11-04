package main

import "fmt"

func media(numeros ...float64) float64{
  total := 0.0
  for _, num := range numeros{
    total += num
  }
  return total/float64(len(numeros))
}

func main(){

  nota := media(7.6,7.8,2.5,5.8,5.9,5.9)
  fmt.Println(nota)

}
