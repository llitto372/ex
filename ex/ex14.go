package main

import (
	"fmt"
)

func appendValue(slice *[]int, value int) {
	*slice = append(*slice, value)
}
func main() {
	nums := []int{1, 2, 3}
	fmt.Println(nums)
	appendValue(&nums, 4)
	fmt.Println(nums)
	appendValue(&nums, 9)
	fmt.Println(nums)
}
