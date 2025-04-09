package main

import (
	"fmt"
	"math"
)

type (
	Point   struct{ x, y float64 } // Point and struct{ x, y float64 } are different types
	polar   Point                  // polar and Point denote different types
	myPoint = Point
)

func (p Point) getDistance(point Point) float64 {
	return math.Sqrt(math.Pow(p.x-point.x, 2) + math.Pow(p.y-point.y, 2))
}

func squareByPoints(p1 Point, p2 Point) float64 {
	return math.Abs((p2.x - p1.x) * (p2.y - p1.y))
}

func square(p1 struct{ x, y float64 }, p2 struct{ x, y float64 }) float64 {
	return math.Abs((p2.x - p1.x) * (p2.y - p1.y))
}

type TreeNode struct {
	left, right *TreeNode
	value       float64
}

func (node TreeNode) squareValue() float64 {
	return math.Pow(node.value, 2)
}

type MyTreeNode TreeNode

func (node MyTreeNode) squareValue() float64 {
	return math.Pow(node.value, 2)
}

// Deposit
/*
Type definition can be used to extend other type
*/
type Deposit int

func (d Deposit) inc() Deposit {
	return d + 1
}

func main() {
	p1 := Point{x: 0, y: 0}
	p2 := myPoint{x: 3, y: 4}
	p1.getDistance(p2)              // syntax sugar for Point.getDistance(p1, p2)
	fmt.Println(p1.getDistance(p2)) // works with Point type and aliases

	fmt.Println("--------------")

	p3 := p1
	p4 := p2
	fmt.Println(squareByPoints(p3, p4)) // works with Point type and aliases
	p5 := polar{3, 4}
	fmt.Println(square(p3, p5)) // works with any type and aliases with underlying type struct{ x, y float64 }

	fmt.Println("--------------")

	treeNode := TreeNode{value: 2}
	fmt.Println(treeNode.squareValue())
	myTreeNode := MyTreeNode{value: 2}
	fmt.Println(myTreeNode.squareValue())

	fmt.Println("--------------")

	deposit := Deposit(1)
	fmt.Println(deposit.inc())
}
