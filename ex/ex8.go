package main

import "fmt"

func sumToN(n int) int {
	var sum int
	for i := 0; i <= n; i++ {
		sum = sum + i
		fmt.Println(sum)
	}
	return sum
}

func main() {
	fmt.Println(sumToN(5))
}
