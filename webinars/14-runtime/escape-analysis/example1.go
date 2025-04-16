package main

import "fmt"

func main() {
	n := 4
	sq := square(n)
	fmt.Println(sq)
}

func square(x int) int {
	return x * x
}
