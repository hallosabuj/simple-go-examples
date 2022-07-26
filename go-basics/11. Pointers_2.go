package main

import "fmt"

type myStruct struct {
	value int
}

func main() {
	var ms *myStruct
	fmt.Println(ms) // <nil>
	ms = &myStruct{value: 10}
	fmt.Println(ms)    // &{10}
	ms = new(myStruct) // Initialize the structure fields with zero values
	fmt.Println(ms)    // &{0}
	*&ms.value = 50
	fmt.Println(ms) // &{50}
	(*ms).value = 100
	fmt.Println(ms) // &{100}
	ms.value = 200
	fmt.Println(ms) // &{200}
}
