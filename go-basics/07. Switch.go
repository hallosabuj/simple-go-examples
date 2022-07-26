package main

import "fmt"

func main() {
	var choice int
	fmt.Printf("Enter  your choice : ")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		fmt.Println("First choice")
	case 2:
		fmt.Println("Second choice")
	case 3:
		fmt.Println("Third choice")
	case 4:
		fmt.Println("Fourth choice")
	default:
		fmt.Println("Default choice")
	}

	switch choice {
	case 1, 11, 21:
		fmt.Println("First choice")
	case 2, 12, 22:
		fmt.Println("Second choice")
	case 3, 13, 23:
		fmt.Println("Third choice")
	case 4, 14, 24:
		fmt.Println("Fourth choice")
	default:
		fmt.Println("Default choice")
	}

	i := 1
	switch i {
	case 1:
		fmt.Println("1")
		fallthrough // This will also execute next case if the current case condition is satisfied
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("Wrong number")
	}
}
