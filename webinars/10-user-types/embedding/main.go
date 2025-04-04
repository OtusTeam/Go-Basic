package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h Human) Walk() {
	fmt.Println("I'm walking")
}

type Student struct {
	Human
	school string
}

func GetAge(human Human) int {
	return human.age
}

func main() {
	student := Student{
		Human{
			"John",
			20,
		},
		"Some school",
	}

	fmt.Println(student.age) // fmt.Println(student.Human.age)
	student.Walk()           // student.Human.Walk()

	//fmt.Println(GetAge(student)) // doesn't compile
	human := Human{
		"John",
		20,
	}
	fmt.Println(GetAge(human))
}
