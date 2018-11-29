package main

import (
	"fmt"
	"time"
)

func routine(c chan int) {
	time.Sleep(time.Second)
	c <- 1 //Operação bloqueante
	fmt.Println("Só depois que foi lido")
}

func main() {
	c := make(chan int) // canal sem buffer

	go routine(c)

	fmt.Println(<-c) //Operação bloqueante
	fmt.Println("Foi lido")
	fmt.Println(<-c) // deadlock
	fmt.Println("Fim")
}
