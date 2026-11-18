package main

import (
	"fmt"
	"strings"
)

func splitwords(text string) []string {
	var words []string
	var current_word string
	//inword := true

	for _, char := range text {
		if char == ' ' || char == '　' {
			if current_word != "" {
				words = append(words, current_word)
				current_word = ""
			}
		} else {
			current_word += string(char)
		}
	}
	if current_word != "" {
		words = append(words, current_word)
	}
	return words
}

func WordCount(s string) map[string]int {

	words := splitwords(s)

	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

func compareMaps(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for key := range a {
		if a[key] != b[key] {
			return false
		}
	}
	return true
}

func WordCountTest(wc func(s string) map[string]int, text string) {
	wordcount := wc(text)

	words := strings.Fields(text)

	trustwc := make(map[string]int)
	for _, word := range words {
		trustwc[word]++
	}
	if compareMaps(trustwc, wordcount) {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
}

func main() {
	text := "I am apple I am Pen youkan is kawaii pupulie is kawaii"
	text2 := "猫と　ああああ　と　わわっわあw wwwww 8888"

	WordCountTest(WordCount, text)
	WordCountTest(WordCount, text2)
}
