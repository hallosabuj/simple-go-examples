package main

import "fmt"

func myFunc(ch chan int) {
	// Send 5 to channel
	ch <- 5
}
func main() {
	fmt.Println("start main thread")
	ch := make(chan int)
	go myFunc(ch)
	valueFromMyFunc := <-ch
	fmt.Println(valueFromMyFunc)
}
