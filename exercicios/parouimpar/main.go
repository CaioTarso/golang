package main

import "fmt"

func main() {
	var n int
	fmt.Println("Digite um número:")
	fmt.Scan(&n)
	fmt.Println(ParOuImpar(n))
}


func ParOuImpar(n int) string {
	if n%2 == 0 {
		return "Par"
	}else {
		return "Ímpar"
	}
}