package main

import "fmt"

func main() {

	mySlice := []int{1, 3, 5, 7, 9, 11}
	fmt.Printf("%T\n", mySlice)
	fmt.Printf("%T\n", &mySlice)

	fmt.Println(mySlice)
	mySlice2 := &mySlice
	fmt.Println(mySlice2)
}
