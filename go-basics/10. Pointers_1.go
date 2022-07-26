package main

import "fmt"

func main() {
	var a int = 10
	var b *int = &a
	fmt.Println(&a, b) // 0xc0000140e8 0xc0000140e8
	fmt.Println(a, b)  // 10 0xc0000140e8
	fmt.Println(a, *b) // 10 10
	a = 20
	fmt.Println(a, *b) // 20 20
	*b = 30
	fmt.Println(a, *b) // 30 30
}
