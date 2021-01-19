package main

import "fmt"

func main() {
	s := "c"
	switch s {
	case "a":
		fmt.Println("The letter is a")
	case "c":
		fmt.Println("The letter is c")
	case "b":
		fmt.Println("The letter is b")
	default:
		fmt.Println("I don't recognize that letter!!!")
	}

}
