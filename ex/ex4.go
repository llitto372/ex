package main

import (
	"fmt"
)

func isLeapYear(year int) bool {
	var exm bool
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		exm = true // или как подсказала мне уже потом иишка можно не использовать if потому что оно тут избыточно и просто через return year%4 == 0 && year%100 != 0 || year%400 == 0
	}
	return exm
}

func main() {
	fmt.Println(isLeapYear(2001))
	fmt.Println(isLeapYear(2020))
	fmt.Println(isLeapYear(540))
}
