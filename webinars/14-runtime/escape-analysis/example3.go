package main

func main() {
	_ = answer()
}

func answer() *int {
	x := 42
	return &x
}
