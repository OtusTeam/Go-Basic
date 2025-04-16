package main

import (
	"fmt"
	"runtime"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("-----")
}

func main() {
	fmt.Println("GOMAXPROCS: ", runtime.GOMAXPROCS(-1))
	//debug.SetGCPercent(-1)
	fmt.Println("-----")

	var mem runtime.MemStats
	printStats(mem)

	for i := 0; i < 10; i++ {
		_ = make([]byte, 50_000_000)
	}
	printStats(mem)

	for i := 0; i < 10; i++ {
		_ = make([]byte, 100_000_000)
	}
	printStats(mem)
}
