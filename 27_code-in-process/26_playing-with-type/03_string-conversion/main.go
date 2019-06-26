package main

import "fmt"

type mySentence string

func main() {
	var message mySentence = "Hello World!"
	fmt.Println(message)
	fmt.Printf("%T\n", message)
	fmt.Printf("%T\n", string(message))
	fmt.Printf("%T\n", []byte(message))
	fmt.Printf("%T\n", []rune(message))
}
