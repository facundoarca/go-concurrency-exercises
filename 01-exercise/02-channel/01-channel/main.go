package main

import (
	"fmt"
)

func main() {

	ch := make(chan int)

	go func(a, b int) {
		c := a + b
		ch <- c
	}(1, 2)

	val, ok := <-ch
	if ok {
		fmt.Print(val)
	}
	// TODO: get the value computed from goroutine
	// fmt.Printf("computed value %v\n", c)
}
