package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkMutex(b *testing.B) {
	wg := sync.WaitGroup{}
	counter := 0
	mutex := sync.Mutex{}

	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer mutex.Unlock()

			mutex.Lock()
			counter++
		}()
	}
	//fmt.Println(counter)
}

func BenchmarkRWMutex(b *testing.B) {
	wg := sync.WaitGroup{}
	counter := 0
	mutex := sync.RWMutex{}

	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer mutex.Unlock()

			mutex.Lock()
			counter++
		}()
	}
	//fmt.Println(counter)
}

func BenchmarkAtomic(b *testing.B) {
	wg := sync.WaitGroup{}
	var counter int64

	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			atomic.AddInt64(&counter, 1)
		}()
	}
	//fmt.Println(counter)
}
