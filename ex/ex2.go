package main

import (
	"errors"
	"fmt"
)

// func secondMax(slice []int) (int, error) {
// 	if len(slice) < 2 {
// 		return 0, errors.New("Должно быть хотя бы 2 элемента")
// 	}
// 	for max1, max2 := slice[0], slice[1]; max1 > max2;
// 	  {

// 	}
// }

func secondMax(slice []int) (int, error) {
	if len(slice) < 2 {
		return 0, errors.New("должно быть хотя бы 2 элемента")

	}
	max1 := slice[0]
	max2 := slice[1]
	if max1 < max2 {
		max1, max2 = max2, max1
	}
	for _, num := range slice[2:] {
		if num > max1 {
			max2 = max1
			max1 = num
		} else if num > max2 {
			max2 = num
		}
	}
	return max2, nil
}

func main() {
	fmt.Println(secondMax([]int{3, 5}))
	fmt.Println(secondMax([]int{}))
	fmt.Println(secondMax([]int{3, 5, 15, 19, 22}))
	fmt.Println(secondMax([]int{3, 412, 235, 95, 84}))
}
