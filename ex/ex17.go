package main

import (
	"fmt"
)

func safeFunction(val func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Ошибочка: ", r)
		}
	}()
	val()
}

func f() {
	fmt.Println("Работаем...")
	panic("проебались, но где")
}

//	func safeFunction() {
//		defer func() {
//			if r := recover(); r != nil {
//				fmt.Println("Ошибочка: ", r)
//			}
//		}()
//		fmt.Println("Работаем...")
//		panic("проебались, но где")
//	}
func main() {
	safeFunction(f)
}
