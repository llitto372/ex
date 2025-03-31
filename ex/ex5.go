package main

import (
	"fmt"
)

func maxOfThree(a, b, c int) int {
	// 	a := arr[0]
	// 	b := arr[1]
	// 	c := arr[2]
	// 	var res int
	// 	if len(arr) != 3 {
	// 		panic("Должно быть 3 элемента")
	// 	}
	// 	if a > b && a > c {
	// 		res = a
	// 	} else if b > a && b > c {
	// 		res = b
	// 	} else if c > a && c > b {
	// 		res = c
	// 	}

	// 	return res

	max := a

	if max < b {
		max = b
	} else if max < c {
		max = c
	}

	return max

}

func main() {
	fmt.Println(maxOfThree(11, 53, 23))
}
