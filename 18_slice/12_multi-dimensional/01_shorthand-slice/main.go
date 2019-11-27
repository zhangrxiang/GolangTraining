package main

import (
	"fmt"
)

func main() {
	student := []string{}
	students := [][]string{}
	fmt.Println(student)
	fmt.Printf("%T\n", student)
	fmt.Printf("%T\n", students)
	fmt.Println(students)
	fmt.Println(student == nil)
	fmt.Println(&student == nil)
	fmt.Println(&student == &[]string{})
}
