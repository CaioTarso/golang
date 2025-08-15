package main

import "fmt"

func main() {
   idade := 12

   if idade >= 21 {
	  fmt.Println("Você já pode beber nos EUA")
   } else if idade >=18 {
	  fmt.Println("Você já pode beber no Brasil")
   } else {
	 fmt.Println("Você é menor de idade")
   }

   personagens := map[string]int{
         "Goku": 8000,
		 "Vegeta": 9000,
	}

	if value, ok := personagens["Seya"]; ok{
		fmt.Println("Poder: ", value, ok)
	}
}

