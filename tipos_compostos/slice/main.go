package main

import "fmt"

func main() {
	var gavetas []string
	gavetas = append(gavetas, "shorts", "camisetas", "calças", "camisas", "sapatos", "meias")
	fmt.Println(gavetas[:3])
	fmt.Println(gavetas[3:])

	o := []int{1, 2, 3, 4, 5}
	fmt.Println(o[2])
}
