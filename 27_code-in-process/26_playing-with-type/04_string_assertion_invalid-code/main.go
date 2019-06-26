package main

import "fmt"

type mySentence string

func f(message mySentence) {
	var message2 interface{} = message
	fmt.Printf("%T\n", message2)
	panic(map[string]interface{}{
		"message": "error",
		"status":  false,
		"code":    500,
	})
	_, err := fmt.Printf("%T\n", message2.(string)) //panic: interface conversion: interface {} is main.mySentence, not string
	if err != nil {
		fmt.Println(err)
	}

}
func p() {
	if err := recover(); err != nil {
		fmt.Println(err)
		a := err.(map[string]interface{})
		fmt.Println(a["message"])
	}
	fmt.Println("----------------over------------------")
}
func main() {
	defer p()
	var message mySentence = "Hello World!"
	fmt.Println(message)
	fmt.Printf("%T\n", message)
	//fmt.Printf("%T\n", message.(string))//invalid type assertion: message.(string) (non-interface type mySentence on left)
	f(message)
}
