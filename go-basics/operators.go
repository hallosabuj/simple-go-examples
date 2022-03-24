package main

import "fmt"

func main() {
	a := 10
	b := 3
	//Arithmatic operators
	fmt.Println("a + b:", a+b)
	fmt.Println("a - b:", a-b)
	fmt.Println("a * b:", a*b)
	fmt.Println("a / b:", a/b)
	fmt.Println("a % b:", a%b)
	//=============================
	//Bitwise operators
	fmt.Println("a & b:", a&b)
	fmt.Println("a | b:", a|b)
	fmt.Println("a ^ b:", a^b)
	fmt.Println("a &^ b:", a&^b)
	fmt.Println("a << 2:", a<<2)
	fmt.Println("a >> 2:", a>>2)
}
