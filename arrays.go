package main

import "fmt"

func main() {
	var name [2]string
	name[0] = "pallavi"
	name[1] = "sawale"

	fmt.Println(name)

	var arr1 = [4]int{1, 2, 3, 4}

	fmt.Println(arr1)
	fmt.Print(len(arr1))
}
