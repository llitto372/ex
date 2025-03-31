package main

import "fmt"

func invertMap(m map[string]int) map[int]string {
	res := make(map[int]string)
	for key, val := range m {
		res[val] = key
	}
	return res
}

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 1}
	fmt.Println(invertMap(m))
}
