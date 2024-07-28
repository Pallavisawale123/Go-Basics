package main

import (
	"fmt"
	"sort"
)

func main() {

	var sliceEx = []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Before ", sliceEx)

	sort.Ints(sliceEx)
	fmt.Println("After sorting : ", sliceEx)

	// append the elements in slice
	sliceEx = append(sliceEx, 22)
	sliceEx = append(sliceEx, 44)

	fmt.Println("total elements after append", sliceEx)

	//convert string to slice

	var str = "pallavi"

	strSlice := []rune(str)
	fmt.Println("First character of string ", string(strSlice[0]))

}
