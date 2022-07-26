package main

import "fmt"

func main() {
	population := make(map[string]int)
	population["India"] = 12
	population["China"] = 13
	if value, ok := population["Japan"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("Key not present in the map")
	}
}
