package main

import (
	"fmt"
)

// Function to reverse any sequence in a string slice
func reverse(s []rune, begin, end int) {
	for begin < end {
		s[begin], s[end] = s[end], s[begin]
		begin++
		end--
	}
}

// Function to reverse words in a string
func reverseWords(s string) string {
	runes := []rune(s)
	wordBegin := -1

	for i := 0; i < len(runes); i++ {
		// This condition checks if the word begins
		if wordBegin == -1 && runes[i] != ' ' {
			wordBegin = i
		}

		// This condition checks if the word ends
		if wordBegin != -1 && (i+1 == len(runes) || runes[i+1] == ' ') {
			reverse(runes, wordBegin, i)
			wordBegin = -1
		}
	}

	// Reverse the entire string
	reverse(runes, 0, len(runes)-1)

	return string(runes)
}

func main() {
	s := "i like this program very much"

	// Function call
	reversed := reverseWords(s)
	fmt.Println(reversed)
}
