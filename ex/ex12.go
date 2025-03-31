package main

import (
	"fmt"
)

func swap(a, b *int) {
	*a, *b = *b, *a
}
func main() {
	w, x, y, z := 1, 2, 3, 4
	fmt.Println(w, x, y, z)
	swap(&x, &w)
	swap(&y, &z)
	fmt.Println(w, x, y, z)
}
