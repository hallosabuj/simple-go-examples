package main

import (
	"fmt"
)

/*
	There are three scope
		1. Package Lavel
		2. Block Scope
		3. Outside of package (variable name started with uppercase)
*/

//OutsideVariable (started using uppercase) variable can be accessed from outside of the package
var OutsideVariable string = "sabuj"

//Here variable can't be defied using := syntax
var i int = 12

//In this way related variables can be grouped together
var (
	name1 string = "sabuj1"
	name2 string = "sabuj2"
	name3 string = "sabuj3"
)

func main() {
	//Variable 1
	var temp = "Sabuj"
	fmt.Println("Hello", temp)
	//Variable 2
	var age int
	age = 23
	fmt.Println("Age :", age)
	//Variable 3
	var weight = 65.5
	fmt.Println("Weight :", weight)
	//Variable 4
	favColor := "Green"
	fmt.Println("Fav Color :", favColor)
	fmt.Printf("%v	%T\n", favColor, favColor)

	//Printing global variable
	fmt.Println("Global variable :", i)
	i := 15
	fmt.Println("Here value will be local value (shadowing) :", i)

	//Type conversion in go
	var flotingvalue float32 = 12.4
	var j = int(flotingvalue) //float32(), float64() ...1
	fmt.Printf("%v	%T", j, j)
}
