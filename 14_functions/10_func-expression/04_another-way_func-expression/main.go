package main

import "fmt"

func makeGreeter() func() string {
	return func() string {
		return "Hello world!"
	}
}

func main() {
	greet := makeGreeter()
	fmt.Println(greet())
	fmt.Println(makeGreeter()())

	f := func() interface{} {
		return 1
	}()

	fmt.Println(f)
}
