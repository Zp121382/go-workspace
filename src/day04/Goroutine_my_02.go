package main

import (
	"fmt"
	"time"
)

func slowFunc() {
	time.Sleep(time.Second * 2)
	fmt.Println("sleeper() finished!!!")
}

func main() {
	go slowFunc()
	fmt.Println("I am now shown until slowFunc() completes!")
	time.Sleep(time.Second * 3)
}
