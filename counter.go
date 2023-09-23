package main

import (
	"strings"
)

// Counter counts the words in given string.
func Counter(s string) int {
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	counter := 0

	for _, line := range lines {
		words := strings.Split(line, " ")
		for _, word := range words {
			word = strings.TrimSpace(word)
			if len(word) > 0 {
				counter++
			}
		}
	}

	return counter
}
