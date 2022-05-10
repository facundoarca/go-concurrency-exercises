package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(4 * time.Second)
		ch <- "one"
	}()

	// TODO: implement timeout for recv on channel ch
	select {
	case m1 := <-ch:
		fmt.Println(m1)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout")
	}

}
