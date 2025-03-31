package main

import (
	"fmt"
	"strings"
	"unicode"
)

func wordCount(text string) map[string]int {
	textrune := []rune(text)
	for i, val := range textrune {
		if !unicode.IsLetter(val) {
			textrune[i] = ' '
		}
	}
	wordslow := strings.ToLower(string(textrune))
	words := strings.Fields(wordslow)
	wordsmap := make(map[string]int)
	for _, value := range words {
		wordsmap[value]++
	}
	return wordsmap
}

func main() {
	fmt.Println(wordCount("HUI !!! HUI !!! HUI< zhopa, zhopa''"))
}
