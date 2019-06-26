package main

import (
	"fmt"
	"math"
)

func main() {
	up := 34.705945
	down := 34.405945
	fmt.Println(math.Floor(up + 0.5))
	fmt.Println(math.Floor(down + 0.5))

	fmt.Println(math.Pi)
	fmt.Println(math.Round(math.Pi))       //3
	fmt.Println(math.RoundToEven(math.Pi)) // 3
	fmt.Println(math.Ceil(math.Pi))        //4
	fmt.Println(math.Floor(math.Pi))       //3
	fmt.Println(math.E)

	fmt.Println(math.Float64frombits(math.Float64bits(-12) &^ (1 << 63))) //12
	fmt.Println(math.Abs(-12))
	fmt.Println(math.Float64bits(-12))                       //13846317054350589952
	fmt.Println(math.Float64bits(12))                        //4622945017495814144
	fmt.Printf("%b\n", math.Float64bits(12))                 //100000000101000000000000000000000000000000000000000000000000000
	fmt.Printf("%b\n", math.Float64bits(-12))                //1100000000101000000000000000000000000000000000000000000000000000
	fmt.Println(math.Float64frombits(math.Float64bits(-12))) //-12
}
