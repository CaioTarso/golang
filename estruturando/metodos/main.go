package main

import "fmt"

type Pessoa struct {
	Nome string
	idade int
}


func (p Pessoa) Apresentar() {
   fmt.Printf("Olá, eu sou %s e tenho %d anos.\n", p.Nome, p.idade)
}

func main() {
	p1:= Pessoa {
	    Nome: "caio",
		idade: 21,
	}
	p1.Apresentar()
}