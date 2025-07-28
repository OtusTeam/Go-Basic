package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter atomic.Int32
	//var counter1 int64
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Add(1)
			//atomic.AddInt64(&counter1, 1)
		}()
	}

	wg.Wait()
	fmt.Println("counter: ", counter.Load())
	//fmt.Println("counter1: ", counter1)
}
