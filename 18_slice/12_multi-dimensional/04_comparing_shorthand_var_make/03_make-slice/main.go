package main

import (
	"fmt"
)

func main() {
	stu := make([]byte, 35)
	student := make([]string, 35)
	students := make([][]string, 35)
	stu[0] = 'z'
	stu[1] = 'i'
	stu[2] = 'n'
	stu[3] = 'g'
	student[0] = string(stu)
	students[0] = student
	// student = append(student, "Todd")
	fmt.Println(student)
	fmt.Println(students)
}
