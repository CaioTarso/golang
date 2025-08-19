package main

import "fmt"

// Implemente uma função Soma(a int, b int) int que recebe dois inteiros e retorna a soma deles.
// No main, leia dois números digitados pelo usuário e mostre o resultado.

func main() {
	fmt.Println("Soma:")
	fmt.Println(Soma(2,3))
}



func Soma(a,b int) int {
	return a + b
}