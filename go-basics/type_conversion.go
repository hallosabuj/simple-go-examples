package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Type conversion examples...")
	var i int = 10
	fmt.Printf("Value : %v, Type : %T\n", i, i)
	//integer to float32
	j := float32(i)
	fmt.Printf("Value : %v, Type : %T\n", j, j)
	//This will not convert integer to string
	str1 := string(i)
	fmt.Printf("Value : %v, Type : %T\n", str1, str1)
	//To convert to string we need to do this
	str2 := strconv.Itoa(i)
	fmt.Printf("Value : %v, Type : %T\n", str2, str2)
	//Float32 to int
	j = 12.5
	fmt.Printf("Value : %v, Type : %T\n", j, j)
	i = int(j)
	fmt.Printf("Value : %v, Type : %T\n", i, i)
}
