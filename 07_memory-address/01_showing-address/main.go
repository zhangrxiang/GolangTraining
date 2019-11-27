package main

import "fmt"

func main() {

	a := 43

	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a)
	fmt.Printf("%d \n", &a)
	fmt.Printf("%x \n", &a)
	fmt.Printf("%#x \n", &a)
}
