package main

import "fmt"

func reverce(slice string) rev {
	slice := []rune()
	for i := 0, j := len(slice)-1; i < j; i++, j-- {
		slice[i] = slice [j]
	}

	rev := string(slice)
}

func main() {
	fmt.Println(reverce("Яшалава"))
}
