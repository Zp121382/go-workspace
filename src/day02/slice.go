package main

import "fmt"

func main() {
	var cheese = make([]string, 2)
	cheese[0] = "Mariolles"
	cheese[1] = "Epoisses de Bourgogne"
	fmt.Println(cheese)
	//cheese = append(cheese[:1], cheese[1+1:]...)
	fmt.Println(len(cheese))

	var smellyCheese = make([]string, 2)
	copy(smellyCheese, cheese[1:])
	fmt.Println(smellyCheese)

}
