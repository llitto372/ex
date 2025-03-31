package main

import (
	"fmt"
)

func removeAtIndex(slice []int, index int) []int {
	if index > len(slice)-1 {
		return slice
	}
	result1 := slice[:index]
	result2 := slice[index+1:]
	result := append(result1, result2...)
	return result
}
func main() {
	fmt.Println(removeAtIndex([]int{1, 4, 5, 6, 7, 8, 3}, 4))
}
