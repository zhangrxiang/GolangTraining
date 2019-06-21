package main

import "fmt"

func main() {

	myGreeting := make(map[string]interface{})
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."
	myGreeting["zing"] = map[string]interface{}{
		"name": "zing",
		"age":  25,
	}
	fmt.Println(myGreeting)
	fmt.Println(myGreeting["zing"])
	z := myGreeting["zing"].(map[string]interface{})
	fmt.Println(z)
	fmt.Println(z["name"])
	fmt.Printf("%T\n", z)

	zz := map[string]interface{}{
		"name": "zing",
		"age":  25,
	}
	fmt.Println(zz["name"], zz["age"])
}
