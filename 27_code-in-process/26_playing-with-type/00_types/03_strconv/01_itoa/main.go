package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int = 5
	str := "Hello world " + strconv.Itoa(x) // int to ascii
	fmt.Println(str)
	fmt.Printf("%s\n", str)
	fmt.Println([]byte(str))
}
