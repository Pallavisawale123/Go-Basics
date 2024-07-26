package main

import (
	"fmt"
)

// Function that takes a pointer to a string as a parameter
func modifyString(s *string) {
	*s = *s + "Hello, Go!"
}

func main() {
	str := "Original"
	fmt.Println("Before:", str)

	// Pass the address of str to the function
	modifyString(&str)

	fmt.Println("After:", str)
}
