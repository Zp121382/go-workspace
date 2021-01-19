package main

import (
	"fmt"
)

func main() {
	s := "After a backslash, certain single character escapes represent special values\nn is a new line \n\t t is a tab"
	s = s + `dawdwadaw dawd 
	dwad
	 dwada
	 dwa
		dawdawdwa`
	fmt.Println(s)
}
