// generator() -> square() -> print

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func generator(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}()
	return out
}

func square(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- func(n int) int {
				time.Sleep(500 * time.Millisecond)
				return n * n
			}(n):
			case <-done:
				return
			}

		}
	}()
	return out
}

func stringify(done <-chan struct{}, in <-chan int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- fmt.Sprint(n):
			case <-done:
				return
			}
		}
	}()
	return out
}

//Fan in
func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {

	out := make(chan int)
	var wg sync.WaitGroup

	//Function that takes the reads from a single channel and inserts into the shared channel
	merger := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
			out <- n
		}
	}

	//Requires waitGroup to ensure no channel is left unmerged.
	wg.Add(len(cs))

	//Iterate over the individual channels and call merger function
	for _, c := range cs {
		go merger(c)
	}

	//Run a separate routine to wait for the individual channels and then close the shared one
	go func() {
		wg.Wait()
		close(out)
	}()

	//Return the out channel (at this point, still open, this happens b4 the previous code block)
	return out

}

func main() {

	done := make(chan struct{})

	t := time.Now()
	in := generator(done, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18)

	// TODO: fan out square stage to run two instances.
	ch1 := square(done, in)
	ch2 := square(done, in)
	ch3 := square(done, in)

	// TODO: fan in the results of square stages.
	ch := merge(done, ch1, ch2, ch3)

	chs := stringify(done, ch)

	//Terminate execution after processing i values
	i := 7
	for s := range chs {
		if i--; i == 0 {
			close(done)
		}
		fmt.Println(s)
	}

	fmt.Println("Running time: ", time.Since(t))

	time.Sleep(550 * time.Millisecond)                       //Wait until all already running goroutines are terminated.
	fmt.Println("Idle goroutines: ", runtime.NumGoroutine()) //Expected output is 1 -> main goroutine.
}
