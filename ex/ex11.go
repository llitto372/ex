package main

import (
	"fmt"
	"math"
)

func powerFunc(ext int) func(num int) int {
	return func(num int) int {
		extf := float64(ext)
		numf := float64(num)
		res := math.Pow(numf, extf)
		return int(res)
	}
}
func main() {
	square := powerFunc(2)
	cube := powerFunc(3)
	fmt.Println(square(10))
	fmt.Println(cube(10))
	fmt.Println(square(6))
	fmt.Println(cube(11))
}
