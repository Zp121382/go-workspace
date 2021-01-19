package main

import "fmt"

type Spuerhero struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	Number int
	Street string
	City   string
}

func main() {
	m := Spuerhero{
		Name: "Batman",
		Age:  20,
		Address: Address{
			Number: 1007,
			Street: "Mountain Drive",
			City:   "Gotham",
		},
	}
	fmt.Println("%+v\n", m)
}
