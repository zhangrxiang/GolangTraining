package main

import "fmt"

func main() {
	var answer int
	var answer2 float64
	answer2 = 32 / 3.74
	fmt.Println(answer2)
	answer3 := 32 / 3.74
	answer = int(answer3)
	fmt.Println(answer)
}
