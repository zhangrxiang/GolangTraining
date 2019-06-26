package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := person{"James", 20}
	p2 := new(person)
	p3 := &person{"James", 20}
	fmt.Println(p1)
	fmt.Println(p1.name)
	fmt.Println(p1.age)
	fmt.Printf("%T\n", p1)
	p2.name = `zing`
	fmt.Println(p2.name)
	fmt.Println(p3.name)

	fmt.Println(p3) //&{James 20}
	fmt.Println(p2) //&{zing 0}
	fmt.Println(p1) //{James 20}

}

// we've already seen the above code
