package main

import (
	"fmt"
	"time"
)

func mensagem() {
	fmt.Println("essa é uma go routine!")
}


func main() {
	go mensagem()
	time.Sleep(1 * time.Second)
	fmt.Println("Isso é a main function")
}