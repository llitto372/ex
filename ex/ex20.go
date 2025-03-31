package main

import "fmt"

func mergeMaps(m1, m2 map[string]int) map[string]int {
	res := make(map[string]int)
	for key1, val1 := range m1 {
		if _, ex := m2[key1]; ex {
			res[key1] = val1 + m2[key1]
		} else {
			res[key1] = val1
		}
	}
	for key2, val2 := range m2 {
		if _, ex := res[key2]; !ex {
			res[key2] = val2
		}
	}
	return res
}

func main() {
	map1 := map[string]int{"apple": 5, "banana": 3}
	map2 := map[string]int{"banana": 2, "orange": 4}
	fmt.Println(mergeMaps(map1, map2))
}
