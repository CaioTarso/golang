package main



import "fmt"

func main() {
	var pessoas = map[string]int{}
	pessoas ["Ray Dalio"] = 70
	pessoas ["Elon Musk"] = 52
	pessoas ["Jeff Bezos"] = 59
	pessoas ["Warren Buffet"] = 92
	fmt.Println(pessoas["Elon Musk"])

	delete(pessoas, "Elon Musk")
	fmt.Println(pessoas)
	
	if idade, existe := pessoas["Elon Musk"]; existe {
		fmt.Println("Idade do Elon Musk:", idade)
		} else {
			fmt.Println("Elon Musk não encontrado")
		}
		
	var capitais = map[string]string{
    "Brasil": "Brasilia",
    "Japão": "Tokyo",
	}
	fmt.Println(capitais)
}