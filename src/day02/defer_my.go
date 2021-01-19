package main

import "fmt"

func main() {
	defer fmt.Println("Test DEFER!!!")
	fmt.Println("After Test DEFFER in defer_my.go")
}
