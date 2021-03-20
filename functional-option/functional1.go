package main

import (
	"fmt"
)

type Car struct {
	make   string
	model  string
	colour string
}

func NewCar() *Car {
	return &Car{
		make:   "BMW",
		model:  "3.18",
		colour: "Silver",
	}
}

func main() {
	car := NewCar()
	fmt.Printf("%+v\n", car)
}
