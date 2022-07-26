package main

import "fmt"

func test1() {
	fmt.Print("Start ")
	defer fmt.Print("Middle ")
	fmt.Print("End ")
}

func test2() {
	defer fmt.Print("Start ")
	defer fmt.Print("Middle ")
	defer fmt.Print("End ")
}

func test3() {
	// In panic situation all the statement before panic will be executed
	// All the deferred statements will be executed in panic situation
	fmt.Print("Start ")
	defer fmt.Print("Middle ")
	panic("Something happened")
	fmt.Print("End ")
}

func test4() {
	// Recover from the panic
	// This will recovr from the panic
	// Current function will not continue but the higher functions will
	fmt.Print("About to panic ")
	defer func() {
		if err := recover(); err != nil {
			fmt.Print(err)
		}
	}()
	panic("Panic haappened ")
	fmt.Print("Panic occurred")
}

func main() {
	// defer keyword defer the execution to the end of the function but before the return statement
	// defer keyword works in a LIFO order
	test1() // Output: Start End Middle
	fmt.Println()
	test2() // Output: End Middle Start
	fmt.Println()
	test4() // About to panic Panic haappened
	fmt.Println()
	test3() // Output: Start Middle panic: Something happened
}
