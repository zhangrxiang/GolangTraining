package main

import "fmt"

func main() {

	var myGreeting map[string]string
	fmt.Println(myGreeting)
	fmt.Println(myGreeting == nil)
	fmt.Println(map[string]string{} == nil) // false
	//fmt.Println(nil == nil) //invalid operation: nil == nil (operator == not defined on nil)
}

// add these lines:
/*
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."
*/
// and you will get this:
// panic: assignment to entry in nil map
