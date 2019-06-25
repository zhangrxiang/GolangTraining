package main

import "fmt"

func main() {
	greeting := "Hello"
	fmt.Println(greeting)
	fmt.Println(greeting[:4])
	fmt.Println(greeting[0:4])
	fmt.Println(greeting[1:4])
	fmt.Println(greeting[1:])
	fmt.Println(greeting[:-2]) // invalid slice index -2 (index must be non-negative)
}
