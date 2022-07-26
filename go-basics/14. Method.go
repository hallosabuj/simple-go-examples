package main

import "fmt"

type Greeter struct {
	name string
	age  int
}

func (g Greeter) Greet() { // Value is used as context
	fmt.Println("Hello ", g.name, " you are ", g.age, " years old")
}
func (g *Greeter) ChangeName() { // Pointer is used as context
	g.name = "John"
}
func main() {
	var g Greeter = Greeter{name: "Jack", age: 30}
	g.Greet() // Hello  Jack  you are  30  years old
	g.ChangeName()
	g.Greet() // Hello  John  you are  30  years old
}
