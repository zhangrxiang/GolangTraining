package main

import "fmt"

func main() {

	mySlice := make([]int, 0, 3)

	fmt.Println("-----------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("-----------------")

	caps := make([]int, 0)
	for i := 0; i < 80000; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len:", len(mySlice), "Capacity:", cap(mySlice), "Value: ", mySlice[i])
		caps = append(caps, cap(mySlice))
	}

	fmt.Println(caps)
	fmt.Println(removeDuplicateElement(caps))

}
func removeDuplicateElement(addrs []int) []int {
	result := make([]int, 0, len(addrs))
	temp := map[int]struct{}{}
	for _, item := range addrs {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
