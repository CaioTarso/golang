package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)

	go func() {
		ch <- 10
		ch <- 20
		ch <- 30
	}()
   
	time.Sleep(1 * time.Second)
	<- ch // Recebe o valor do canal
	<- ch // Recebe o próximo valor do canal
	valor := <-ch
	fmt.Println("Valor recebido do canal:", valor)
}
