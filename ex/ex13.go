package main

import (
	"fmt"
)

func incr(ext *int) int {
	*ext = *ext + 1
	return *ext
}
func main() {
	x := 5
	fmt.Println(x)
	incr(&x)
	fmt.Println(x)
}
