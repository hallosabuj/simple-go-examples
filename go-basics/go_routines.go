package main

import "fmt"

func test() {
	for i := 0; i < 100; i++ {
		fmt.Println("Inside test :", i)
	}
}
func main() {
	fmt.Println("Hello inside main function")
	//To call a function in another thread other than the main thread
	go test()
	for i := 0; i < 100; i++ {
		fmt.Println("Inside main :", i)
	}
}
