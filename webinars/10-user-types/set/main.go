package main

import "fmt"

type IntSet map[int]struct{}

func main() {
	ints := []int{1, 2, 3, 7, 3, 2}
	intSet := make(IntSet)
	for _, i := range ints {
		intSet[i] = struct{}{}
	}

	fmt.Println(intSet)
}
