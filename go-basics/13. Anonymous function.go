package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello I am an anonymous function") // Hello I am an anonymous function
	}()

	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Print(i, " ") // 0 1 2 3 4
		}(i)
	}
	fmt.Println()

	var f func(num int) = func(number int) {
		fmt.Println("Passed number is : ", number) // Passed number is :  10
	}
	f(10)
}
