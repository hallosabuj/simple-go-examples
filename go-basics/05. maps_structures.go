package main

import "fmt"

type Student struct {
	Name  string
	Marks int
}

func main() {
	fmt.Println("Maps")
	// We can't use map only after declaring we need initialize the map object
	statePopulation := map[string]int{
		"India": 10,
		"China": 11,
	}
	fmt.Println(statePopulation) // map[China:11 India:10]
	ages := make(map[string]int, 1)
	ages["Jack"] = 20
	ages["Jill"] = 25
	ages["Jhon"] = 25
	fmt.Println(ages) // map[Jack:20 Jhon:25 Jill:25]
	delete(ages, "Jill")
	fmt.Println(ages) // map[Jack:20 Jhon:25]
	value, ok := ages["Jill"]
	fmt.Println(value, ok) // 0 false because Jill is not present in the map
	fmt.Println(len(ages)) // 2

	fmt.Println("Struct")
	s1 := Student{Name: "John", Marks: 30} // {John 30}
	fmt.Println(s1)
	s2 := Student{"Smita", 15}
	fmt.Println(s2) // {Smita 15}

	s3 := struct {
		Name  string
		Marks int
	}{"Smith", 15}
	fmt.Println(s3) // {Smith 15}

	// Structure embeding
	type Animal struct {
		Name string
	}
	type Bird struct {
		Animal
		CanFly bool
	}
	b := Bird{}
	b.Name = "Dove"
	b.CanFly = true
	fmt.Println(b)
}
