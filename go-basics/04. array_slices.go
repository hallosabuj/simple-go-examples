package main

import "fmt"

func main() {
	// this is slice
	grades := []int{1, 2, 3}        // In sloce we can append data but in array we can't append new data
	fmt.Println("Grades :", grades) // Grades : [1 2 3]
	// this is array
	marks := [...]int{1, 2, 4}
	fmt.Println("Marks :", marks) // Marks : [1 2 4]
	ages := [3]int{10, 11, 12}
	fmt.Println("Ages :", ages) // Ages : [10 11 12]

	var students [3]string
	students[0] = "sabuj"
	fmt.Println("Students :", students)           // Students : [sabuj  ]
	fmt.Println("# of students :", len(students)) // # of students : 3

	// Identity matrix
	var identityMatrix [3][3]int = [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	var matrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}} // Here redundant type declared
	fmt.Println("Indentity matrix :", identityMatrix)                                   //Indentity matrix : [[1 0 0] [0 1 0] [0 0 1]]
	fmt.Println("Matrix :", matrix)                                                     // Matrix : [[1 0 0] [0 1 0] [0 0 1]]

	// For array
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 10             // only change value in b
	fmt.Println("a :", a) //a : [1 2 3]
	fmt.Println("b :", b) //b : [1 10 3]
	c := &a
	c[1] = 12             // change value in both a and c
	fmt.Println("a :", a) // a : [1 12 3]
	fmt.Println("c :", c) // c : &[1 12 3]

	// For slice
	x := []int{1, 2, 3} // slice is a projection that means itself it works as pointer
	y := x
	y[1] = 15             // change value in both x and y
	fmt.Println("x :", x) // x : [1 15 3]
	fmt.Println("y :", y) // y : [1 15 3]

	//Here all the slices pointing to the same underling array
	aa := []int{1, 2, 3, 4, 5, 6, 7, 8}
	ab := aa[:]
	ac := aa[3:]
	ad := aa[:3]
	ae := aa[3:5]
	fmt.Println(aa) // [1 2 3 4 5 6 7 8]
	fmt.Println(ab) // [1 2 3 4 5 6 7 8]
	fmt.Println(ac) //[4 5 6 7 8]
	fmt.Println(ad) //[1 2 3]
	fmt.Println(ae) //[4 5]

	aaa := make([]int, 3)               // Making an slice of lenght 3
	fmt.Println(aaa)                    // [0 0 0]
	fmt.Println("Length :", len(aaa))   // Length : 3
	fmt.Println("Capacity :", cap(aaa)) // Capacity : 3

	aab := make([]int, 3, 100)          // Making an slice of length 3 and capacity 100
	fmt.Println(aab)                    // [0 0 0]
	fmt.Println("Length :", len(aab))   // Length : 3
	fmt.Println("Capacity :", cap(aab)) // Capacity : 100

	aac := []int{}
	fmt.Println("Length :", len(aac))   // Length : 0
	fmt.Println("Capacity :", cap(aac)) // Capacity : 0
	aac = append(aac, 1)
	fmt.Println("Length :", len(aac))   // Length : 1
	fmt.Println("Capacity :", cap(aac)) // Capacity : 1
	aac = append(aac, aab...)
	fmt.Println("Length :", len(aac))   // Length : 4
	fmt.Println("Capacity :", cap(aac)) // Capacity : 4
}
