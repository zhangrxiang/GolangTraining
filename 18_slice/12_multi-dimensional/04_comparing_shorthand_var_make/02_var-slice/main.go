package main

import (
	"fmt"
)

func main() {
	var student []string
	var students [][]string

	fmt.Println(student)        //[]
	fmt.Println(student == nil) //true
	fmt.Println(students)       //[]

	//panic: runtime error: index out of range
	student[0] = "Todd"

	// student = append(student, "Todd")
	fmt.Println(student)
	fmt.Println(students)
}
