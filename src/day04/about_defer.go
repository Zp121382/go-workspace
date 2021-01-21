package main

import "fmt"

func main() {
	var i int = 1
	defer fmt.Println("result1 =>", func() int { return i * 2 }())
	i++

	defer func() {
		fmt.Println("result2 =>", i*2)
	}()
	i++

}
