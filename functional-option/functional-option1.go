package main

import (
	"fmt"
)

type Car struct {
	make   string
	model  string
	colour string
}

type Option func(*Car)

func Make(v string) Option {
	return func(c *Car) {
		c.make = v
	}
}
func Model(v string) Option {
	return func(c *Car) {
		c.model = v
	}
}
func Colour(v string) Option {
	return func(c *Car) {
		c.colour = v
	}
}

func NewCar(opts ...Option) *Car {
	c := &Car{
		make:   "BMW",
		model:  "3.18",
		colour: "Silver",
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func main() {
	car1 := NewCar()
	fmt.Printf("%+v\n", car1)

	car2 := NewCar(Make("Mercedes"), Model("SEL 500"))
	fmt.Printf("%+v\n", car2)
}
