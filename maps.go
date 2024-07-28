package main

import (
	"fmt"
)

func main() {

	mapp := make(map[string]int)
	mapp["pallvi"] = 25
	mapp["pooja"] = 26

	fmt.Println("mapp contains :", mapp)

	mapp["rosh"] = 24
	fmt.Println("mapp contains :", mapp)

	delete(mapp, "rosh")

	fmt.Println("mapp contains :", mapp)

	// map iteration
	for k, v := range mapp {
		fmt.Printf("mapp contains key %v and value %v \n ", k, v)
	}
}
