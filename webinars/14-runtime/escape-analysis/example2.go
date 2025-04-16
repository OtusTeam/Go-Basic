package main

import "fmt"

func main() {
	n := 4
	squareByPointer(&n)
	fmt.Println(n)
}

func squareByPointer(x *int) {
	*x = *x * *x
}
