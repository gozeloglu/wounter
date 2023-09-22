package main

import (
	"strings"
)

// Counter counts the words in given string.
func Counter(s string) int {
	s = strings.TrimSpace(s)
	words := strings.Split(s, " ")
	counter := 0

	for _, word := range words {
		word = strings.TrimSpace(word)
		if len(word) > 0 {
			counter++
		}
	}

	return counter
}
