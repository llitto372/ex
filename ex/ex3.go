package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(mergeSortedArrays([]int{1, 3, 5}, []int{2, 4, 6})) // [1 2 3 4 5 6]
}

func mergeSortedArrays(s1, s2 []int) []int {
	slice := make([]int, 0)

	for _, value := range s1 { // [1, 3, 5]
		slice = append(slice, value)
	}

	for _, value := range s2 { // [2, 4, 6]
		slice = append(slice, value)
	}

	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	return slice

}
