package main

import "fmt"

func main() {
	// strrev := rev("pallavi")
	// fmt.Println(strrev)
	revEachWords("pallavi")
}

func rev(name string) (result string) {

	for _, v := range name {
		fmt.Println(string(v))
		result = string(v) + result
	}
	return
}

func revEachWords(word string) {
	s := []rune(word)
	fmt.Println(string(s))

	begin := 0
	end := len(s) - 1

	for begin < end {
		s[begin], s[end] = s[end], s[begin]
		begin++
		end--
	}

	fmt.Println(string(s))
}
