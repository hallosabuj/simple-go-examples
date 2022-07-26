package main

import "fmt"

func main() {
	var i interface{} = 0
	// Type switching
	switch i.(type) {
	case int:
		fmt.Println("i is integer")
	case string:
		fmt.Println("i is string")
	default:
		fmt.Println("i is of unknown type")
	}
}
