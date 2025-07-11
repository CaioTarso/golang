package newpacote

import "fmt"

//As funções que eu quero exportar, devem obrigatoriamente comecarem com a letra maiuscula
//Enquanto as que eu quero que sejam executadas apenas no meu modulo, eu uso minusculo
func SayHi() {
	fmt.Println("Say HI")
}