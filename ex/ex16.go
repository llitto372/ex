package main

import (
	"fmt"
)

func SafeDivide(a, b int) int {
	defer func() {
		if r := recover(); r != nil {
			// fmt.Println("0")
		}
	}()
	res := a / b
	if b == 0 {
		panic("b=0")
	}
	return res
}
func main() {
	fmt.Println(SafeDivide(239, 7))
	fmt.Println(SafeDivide(2, 2))
	fmt.Println(SafeDivide(2345, 0))
	fmt.Println("РАБОТАЕМ БРАТЬЯ")
}
