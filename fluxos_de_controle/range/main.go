package main

import "fmt"

func main () {

	users := map[string]int{
		"caiotarso": 30,
		"joao":      25,
		"maria":     28,
	}

	for key, _ := range users {
		fmt.Println(key)
	}
}
