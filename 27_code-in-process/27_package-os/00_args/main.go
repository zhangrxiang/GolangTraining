package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) > 1 {
		entry := os.Args[1]
		fmt.Println(os.Args[0])
		fmt.Println(entry)
	} else {
		fmt.Println(fmt.Println(os.Args))
	}

}

/*
step 1:
go install

step 2 - from terminal:
programName arguments

*/
