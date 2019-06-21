package main

import "fmt"

func main() {
	letter := rune("A"[0])
	fmt.Println(letter)

	letter = rune(65)
	fmt.Println(letter)

	letter = rune('A')
	fmt.Println(letter)
}
