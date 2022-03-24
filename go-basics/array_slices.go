package main

import "fmt"

func main() {
	// this is slice
	grades := []int{1, 2, 3}
	fmt.Println("Grades :", grades)
	// this is array
	marks := [...]int{1, 2, 4}
	fmt.Println("Marks :", marks)
	ages := [3]int{10, 11, 12}
	fmt.Println("Ages :", ages)

	var students [3]string
	students[0] = "sabuj"
	fmt.Println("Students :", students)
	fmt.Println("# of students :", len(students))

	// Identity matrix
	var identityMatrix [3][3]int = [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	var matrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}} // Here redundant type declared
	fmt.Println("Indentity matrix :", identityMatrix)
	fmt.Println("Matrix :", matrix)

	// For array
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 10 // only change value in b
	fmt.Println("a :", a)
	fmt.Println("b :", b)
	c := &a
	c[1] = 12 // change value in both a and c
	fmt.Println("a :", a)
	fmt.Println("c :", c)

	// For slice
	x := []int{1, 2, 3} // slice is a projection that means itself it works as pointer
	y := x
	y[1] = 15 // change value in both x and y
	fmt.Println("x :", x)
	fmt.Println("y :", y)

	//Here all the slices pointing to the same underling array
	aa := []int{1, 2, 3, 4, 5, 6, 7, 8}
	ab := aa[:]
	ac := aa[3:]
	ad := aa[:3]
	ae := aa[3:5]
	fmt.Println(aa)
	fmt.Println(ab)
	fmt.Println(ac)
	fmt.Println(ad)
	fmt.Println(ae)

	aaa := make([]int, 3) // Making an slice of lenght 3
	fmt.Println(aaa)
	fmt.Println("Length :", len(aaa))
	fmt.Println("Capacity :", cap(aaa))

	aab := make([]int, 3, 100) // Making an slice of length 3 and capacity 100
	fmt.Println(aab)
	fmt.Println("Length :", len(aab))
	fmt.Println("Capacity :", cap(aab))

	aac := []int{}
	fmt.Println("Length :", len(aac))
	fmt.Println("Capacity :", cap(aac))
	aac = append(aac, 1)
	fmt.Println("Length :", len(aac))
	fmt.Println("Capacity :", cap(aac))
	aac = append(aac, aab...)
	fmt.Println("Length :", len(aac))
	fmt.Println("Capacity :", cap(aac))
}
