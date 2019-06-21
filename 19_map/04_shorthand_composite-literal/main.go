package main

import "fmt"

func main() {

	myGreeting := map[string]string{}
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting)

	for key, value := range myGreeting {
		fmt.Println(key, value)
	}
}
