package main

import (
	"fmt"
)

func anotherfunction(f func() string) string {
	return f()
}

func main() {
	fn := func() string {
		return "function test"
	}
	fmt.Println(anotherfunction(fn))
}
