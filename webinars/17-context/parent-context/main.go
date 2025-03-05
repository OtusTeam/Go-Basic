package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func runForever(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context cancelled, err: ", ctx.Err())
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("Running")
		}
	}
}

func main() {
	waitGroup := sync.WaitGroup{}
	background := context.Background()
	withCancel, cancelFunc := context.WithCancel(background)
	defer cancelFunc()
	withTimeout, c := context.WithTimeout(withCancel, 5*time.Second)
	defer c()

	waitGroup.Add(1)
	go runForever(withTimeout, &waitGroup)
	time.Sleep(3 * time.Second)
	cancelFunc()
	waitGroup.Wait()
}
