package main

import (
	"fmt"
)

func sequenceGenerator() func() int {
	i := 1
	return func() int {
		res := i
		i++
		return res
	}

}

func main() {
	next := sequenceGenerator()
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
}
