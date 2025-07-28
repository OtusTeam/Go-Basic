package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func getSomething() int {
	var result int

	once.Do(
		func() {
			fmt.Println("sync.Once")
			result++
		},
	)

	return result
}

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(getSomething())
		}()
	}

	wg.Wait()
}
