package main

import "fmt"

func main() {
	n := answer()
	fmt.Println(*n)
}

func answer() *int {
	x := 42
	return &x
}
