package main

import (
	"fmt"
	"strconv"
)

func reverceNumber(n int) int {
	if n < 0 {
		panic("Введи положительное число")
	}
	num := strconv.Itoa(n)
	runes := []rune(num)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	res := string(runes)
	res1, err := strconv.Atoi(res)
	if err != nil {
		fmt.Println("HUI", err)
	}
	return res1
}

func main() {
	fmt.Println(reverceNumber(245762345))
}
