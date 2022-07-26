package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	} // Output: 0 1 2 3 4
	fmt.Println()

	i := 0
	for i < 5 {
		fmt.Print(i, " ")
		i++
	} // Output: 0 1 2 3 4
	fmt.Println()

	j := 0
	for {
		if j >= 5 {
			break
		}
		fmt.Print(j, " ")
		j++
	} // Output: 0 1 2 3 4
	fmt.Println()
}
