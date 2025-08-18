package main

import "fmt"

func main() {
	soma(10, 20)
	subtracao(90, 20)

	multiplica := func(x int) int {
		return x * 2
	}

	resultado := multiplica(4)
	fmt.Println(resultado)
}


func soma(a, b int) {
	fmt.Println(a+b)
}

func subtracao(a, b int) {
	fmt.Println(a-b)
}
