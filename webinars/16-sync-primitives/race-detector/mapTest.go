package main

import (
	"fmt"
	"sync"
)

type IntSet map[int]struct{}

func main() {
	ints := IntSet{}
	cycles := 100
	waitGroup := sync.WaitGroup{}

	for i := 0; i < cycles; i++ {
		waitGroup.Add(1)
		go func(i int) {
			defer waitGroup.Done()
			ints[i] = struct{}{}
		}(i)
	}

	waitGroup.Wait()

	for i := range ints {
		fmt.Println(i)
	}
}
