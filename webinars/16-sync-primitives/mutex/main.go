package main

import (
	"fmt"
	"sync"
)

func main() {
	//runtime.GOMAXPROCS(1)

	const count = 100
	wg := sync.WaitGroup{}
	v := 0
	var mu sync.RWMutex
	//mu.RLocker()

	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			defer mu.Unlock()

			mu.Lock()
			v++
		}()
	}

	wg.Wait()
	fmt.Println(v)
}
