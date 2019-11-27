package main

import "fmt"

const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
	TB = 1 << (iota * 10) // 1 << (4 * 10)
)

const (
	_   = iota             // 0
	KB2 = 1 << (iota * 10) // 1 << (1 * 10)
	MB2                    // 1 << (2 * 10)
	GB2                    // 1 << (3 * 10)
	TB2                    // 1 << (4 * 10)
)

func main() {
	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b %b\t", KB, KB2)
	fmt.Printf("%d %d\n", KB, KB2)
	fmt.Printf("%b %b\t", MB, MB2)
	fmt.Printf("%d %d\n", MB, MB2)
	fmt.Printf("%b %b\t", GB, GB2)
	fmt.Printf("%d %d\n", GB, GB2)
	fmt.Printf("%b %d\t", TB, TB2)
	fmt.Printf("%d %d\n", TB, TB2)
}
