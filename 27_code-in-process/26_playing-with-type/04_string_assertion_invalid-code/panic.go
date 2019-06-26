package main

import "fmt"

func main() {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("c")
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
			e := err.(map[string]interface{})
			fmt.Println(e["message"]) //
			fmt.Println(e["status"])  //
			fmt.Println(e["code"])    //
		}
		fmt.Println("d")
	}()
	f2()
	defer f3()
}

func f3() {
	if err := recover(); err != nil {
		fmt.Println(err) // 这里的err其实就是panic传入的内容，55
		e := err.(map[string]interface{})
		fmt.Println(e["message"]) //
		fmt.Println(e["status"])  //
		fmt.Println(e["code"])    //
	}
}
func f2() {
	fmt.Println("a")
	//panic(55)
	panic(map[string]interface{}{
		"message": "error",
		"status":  false,
		"code":    500,
	})
	fmt.Println("b")
	fmt.Println("f")
}
