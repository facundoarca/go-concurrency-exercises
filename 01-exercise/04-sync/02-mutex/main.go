package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	runtime.GOMAXPROCS(4)

	var balance int
	var wg sync.WaitGroup
	var mu sync.RWMutex

	deposit := func(amount int) {
		mu.Lock()
		balance += amount
		mu.Unlock()
	}

	checkBalance := func() int {
		mu.RLock()
		defer mu.RUnlock()
		return balance
	}

	wg.Add(30)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			fmt.Println("Increasing balance by $1")
			deposit(1)
		}()
	}
	for i := 0; i < 20; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println("(", i, ") Value read: ", checkBalance())
		}(i)
	}

	//TODO: implement concurrent read.
	// allow multiple reads, writes holds the lock exclusively.

	wg.Wait()
	fmt.Println(balance)
}
