package main

import (
	"sync"
	"testing"
)

var (
	counters        = NewCounters()
	syncMapCounters = sync.Map{}
	wg              = sync.WaitGroup{}
)

func BenchmarkCounters_Load(b *testing.B) {
	counters.Clear()
	for i := 0; i < 1_000; i++ {
		counters.Store(i, i)
	}

	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_, _ = counters.Load(i)
		}(i)
	}

	wg.Wait()
}

func BenchmarkSyncMap_Load(b *testing.B) {
	syncMapCounters.Clear()
	for i := 0; i < 1_000; i++ {
		syncMapCounters.Store(i, i)
		//syncMapCounters.Store(strconv.Itoa(i), i)
	}

	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_, _ = syncMapCounters.Load(i)
		}(i)
	}

	wg.Wait()
}

func BenchmarkCounters_Store(b *testing.B) {
	counters.Clear()

	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			counters.Store(i, i)
		}(i)
	}

	wg.Wait()
}

func BenchmarkSyncMap_Store(b *testing.B) {
	syncMapCounters.Clear()

	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			syncMapCounters.Store(i, i)
		}(i)
	}

	wg.Wait()
}
