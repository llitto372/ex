package main

import (
	"fmt"
	"strings"
	"unicode"
)

type TextProcessor struct {
	Text string
}

func (t *TextProcessor) WordCount() map[string]int {
	runetext := []rune(t.Text)
	for i, val := range runetext {
		if !unicode.IsLetter(val) {
			runetext[i] = ' '
		}
	}
	newtext := strings.ToLower(string(runetext))
	slicetext := strings.Fields(newtext)
	maptext := make(map[string]int)
	for _, value := range slicetext {
		maptext[value]++
	}
	return maptext
}
func (t *TextProcessor) ReplaceWord(old, new string) {
	t.Text = strings.ReplaceAll(strings.ToLower(t.Text), old, new)
}
func main() {
	tp := TextProcessor{"Hello world, hello again!"}
	fmt.Println(tp.WordCount())
	tp.ReplaceWord("hello", "hi")
	fmt.Println(tp.Text)
}
