package main

import "fmt"

type loja struct {
	Nome string
	Endereco string
	Telefone string
	Produtos []produto
}

type produto struct {
	Nome string
	Preco float64
	Quantidade int
}

func main() {
	loja1 := loja{
		Nome: "Loja Exemplo",
		Endereco: "Rua Exemplo, 123",
		Telefone: "1234-5678",
		Produtos: []produto{
			{Nome: "Produto A", Preco: 10.99, Quantidade: 5},
			{Nome: "Produto B", Preco: 20.49, Quantidade: 3},
			{Nome: "Produto C", Preco: 15.75, Quantidade: 10},
},
}
    loja1.Nome = "Auto Peças São Luiz"
	fmt.Println("Loja:", loja1)

}
