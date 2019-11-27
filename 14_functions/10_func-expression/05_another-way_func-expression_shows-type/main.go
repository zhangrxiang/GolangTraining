package main

import "fmt"

func makeGreeter() func() string {
	return func() string {
		return "Hello world!"
	}
}

func makeGreeter2(hello string) string {
	return func(hello string) (hello2 string) {
		hello2 = hello
		return hello2
	}(hello)
}

func main() {
	greet := makeGreeter()
	fmt.Println(greet())
	fmt.Printf("%T\n", greet)
	greeter2 := makeGreeter2("hello world")
	fmt.Println(greeter2)
}
