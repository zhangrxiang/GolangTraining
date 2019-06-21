package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"Tim":   "Good morning!",
		"Jenny": "Bonjour!",
	}

	myGreeting["Harleen"] = "Howdy"

	fmt.Println(len(myGreeting))
	myGreeting["Harleen"] = "Howdy"
	fmt.Println(len(myGreeting)) //3
	myGreeting["zing"] = "Howdy"
	fmt.Println(len(myGreeting)) //4
}
