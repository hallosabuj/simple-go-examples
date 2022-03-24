package main

import "fmt"

func main() {
	//Default value for boolean is 0 which is false
	var flag = false
	fmt.Printf("%v : %T\n", flag, flag)
	var i int8 = 34 //int16, int32, int64
	fmt.Printf("%v : %T\n", i, i)
	var j uint8 = 234 //uint16, uint32
	//byte is similar to uint8
	fmt.Printf("%v : %T\n", j, j)
	var n = 2.0 //By default float64 there is another float32
	n = 3.1e72  //3.1E72
	fmt.Printf("%v : %T\n", n, n)

	var c complex64 = 1 + 2i //complex128
	fmt.Printf("%v : %T\n", c, c)
	c = 1
	fmt.Printf("%v : %T\n", c, c)
	c = 1i
	fmt.Printf("%v : %T\n", c, c)
	c = complex(2, 3)
	fmt.Printf("%v : %T\n", real(c), real(c))
	fmt.Printf("%v : %T\n", imag(c), imag(c))
	////////////////////////////////////////////
	str := "This is string"
	fmt.Printf("%v : %T\n", str, str)
	fmt.Printf("%v : %T\n", str[2], str[2])                 //It will result uint8
	fmt.Printf("%v : %T\n", string(str[2]), string(str[2])) //It will result string
	fmt.Printf("%v : %T\n", []byte(str), []byte(str))       //It will result []uint8
	var r rune = 'a'
	var rr = 'b'
	fmt.Printf("%v : %T\n", r, r)   //It will result int32
	fmt.Printf("%v : %T\n", rr, rr) //It will result int32
}
