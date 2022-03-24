package main

import "fmt"

// This is enumerated constant
// iota is scoped to a constant block
const (
	s = iota
	t = iota
	u = iota
)
const (
	v = iota
)
const (
	w = iota + 10 // 0 + 10
	x             // 1 + 10
	y             // 2 + 10
	z             // 3 + 10
)
const (
	_  = iota
	KB = 1 << (10 * iota) // 000...00001  <<  00000010000000000
	MB                    // 000...00001  <<  00000100000000000
	GB                    // 000...00001  <<  00001000000000000
	TB                    // 000...00001  <<  00010000000000000
	PB                    // 000...00001  <<  00100000000000000
	EB                    // 000...00001  <<  10000000000000000
)

func main() {
	fmt.Println("Here constants are elaorated....")
	// constant can remain unsed through out the program
	const a int = 27
	const b int = 23
	const c = 24
	var d int16 = 15
	// fmt.Printf("%v : %T\n", c+d, c+d) // mismatched types int and int16
	fmt.Printf("%v : %T\n", c, c)
	fmt.Printf("%v : %T\n", c+d, c+d) // this will add successfully
	//iota
	fmt.Println(s)
	fmt.Println(t)
	fmt.Println(u)
	fmt.Println(v)
	fmt.Println("w :", w)
	fmt.Println("x :", x)
	fmt.Println("y :", y)
	fmt.Println("z :", z)
	fmt.Println("KB :", KB)
}
