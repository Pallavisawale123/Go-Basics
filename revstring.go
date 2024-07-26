package main

import "fmt"

func main() {
	strrev := rev("pallavi")
	fmt.Println(strrev)
}

func rev(name string) (result string) {

	for _, v := range name {
		fmt.Println(string(v))
		result = string(v) + result
	}
	return
}
