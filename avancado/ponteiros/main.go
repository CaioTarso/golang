package main

import "fmt"

type Pessoa struct {
	Nome string
}

func main () {
	p1 := Pessoa{Nome: "Caio"}
	p2 := Pessoa{Nome: "Maria"}

    var ponteiro *Pessoa = &p1
	ponteiro.Nome = "João"
	fmt.Println(ponteiro.Nome)

	fmt.Println(p1.Nome)
	fmt.Println(p2.Nome)
}