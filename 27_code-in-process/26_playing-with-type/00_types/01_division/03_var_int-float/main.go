package main

import (
	"fmt"
)

func main() {
	answer := 32 / 3.74
	fmt.Println(answer)      //8.556149732620321
	fmt.Println(int(answer)) //8

	var i interface{} = nil
	i = answer
	fmt.Println(i)
	fmt.Println(i.(float64))
}
