package main

import (
	"GolangTraining/04_scope/01_package-scope/02_visibility/vis"
	"fmt"
)

func main() {
	fmt.Println(vis.MyName)
	vis.PrintVar()
}
