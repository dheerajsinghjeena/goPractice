package main

import (
	"fmt"
	"time"
)

func printNumber() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func main() {
	go printNumber()
	time.Sleep(1 * time.Second)
}
