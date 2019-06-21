package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var myGreeting = make(map[string]string)
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting)
	fmt.Println(json.Marshal(myGreeting))
	bytes, _ := json.Marshal(myGreeting)
	fmt.Printf("%s\n", string(bytes))
	v := make(map[string]string)
	_ = json.Unmarshal(bytes, &v)
	fmt.Println(v)
	fmt.Println(v["Tim"])
}
