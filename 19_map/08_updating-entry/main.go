package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"Tim":   "Good morning!",
		"Jenny": "Bonjour!",
	}

	myGreeting["Harleen"] = "Howdy"
	fmt.Println(myGreeting)
	myGreeting["Harleen"] = "Gidday"
	fmt.Println(myGreeting)
	delete(myGreeting, "Tim")
	fmt.Println(myGreeting)
	delete(myGreeting, "Tim")
	fmt.Println(myGreeting)
}
