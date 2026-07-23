package main

func main() {
	n := 4
	squareByPointer(&n)
}

func squareByPointer(x *int) {
	*x = *x * *x
}
