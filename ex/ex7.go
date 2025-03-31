package main

import "fmt"

func calculator(a, b int, oper string) int {
	var res int
	if b == 0 && oper == "/" {
		panic("На ноль делить нельзя")
	}
	switch oper {
	case "-":
		res = a - b
		return res
	case "+":
		res = a + b
		return res
	case "/":
		res = a / b
		return res
	case "*":
		res = a * b
		return res
	}
	return res
}

func main() {
	fmt.Println(calculator(1000, 7, "-"))
	fmt.Println(calculator(32, 1, "*"))
	fmt.Println(calculator(23, 54, "-"))
	fmt.Println(calculator(3, 5, "*"))
}
