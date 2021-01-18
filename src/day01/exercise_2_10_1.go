package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var s string = "1234"
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))
	var i int
	i, _ = strconv.Atoi(s)
	fmt.Println(i)
	fmt.Println(reflect.TypeOf(i))
}
