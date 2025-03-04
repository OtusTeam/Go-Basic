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
	ctx, cancelFunc := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancelFunc()

	waitGroup.Add(1)
	go runForever(ctx, &waitGroup)
	time.Sleep(3 * time.Second)
	cancelFunc()
	waitGroup.Wait()
}
