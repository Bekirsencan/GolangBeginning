package main


import (
	"fmt"
)

func main() {
	grades := [...]int{97,85,93}
	fmt.Printf("Grades: %v", grades)
}

// Reaching Array's Elements

func reach() {
	var students [3]string
	fmt.Printf("Students: %v", students)
	students[0] = "Lisa"
	fmt.Printf("Students: %v", students)
	fmt.Printf("Number of Students: %v",len(students))
}

// Array of Arrays,

func AOA() {
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1,0,0}
	identityMatrix[1] = [3]int{0,1,0}
	identityMatrix[2] = [3]int{0,0,1}
	fmt.Println(identityMatrix)
}

// Array Copying And Pointing

func copy() {
	a := [...]int{1,2,3}
	b := &a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
}