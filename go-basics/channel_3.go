package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	// anonymous go routine
	go func() {
		ch <- "Message 1"
		ch <- "Message 2"
		ch <- "Message 3"
		close(ch)
	}()

	// Using for loop over the channel
	for value := range ch {
		fmt.Println(value)
	}
}
