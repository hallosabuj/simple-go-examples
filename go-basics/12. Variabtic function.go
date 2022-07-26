package main

import "fmt"

func main() {
	sum(1, 2, 3, 4, 5)
}

func sum(values ...int) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	fmt.Println("Sum is : ", sum)
}
