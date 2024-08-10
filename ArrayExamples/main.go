package main

import "fmt"

func main() {

	var testArray [4]string
	testArray[0] = "pallavi"
	testArray[1] = "pooja"
	testArray[2] = "rosh"
	testArray[3] = "Dipps"
	fmt.Println("Values of testArray :", testArray)

	var testArray1 = []int{1, 2, 3, 4, 5}
	fmt.Println("Values of testArray1 :", testArray1)

	//Print the values of positive index
	var i = 0
	for i < len(testArray) {
		if i%2 == 0 {
			fmt.Println(testArray[i])
		}
		i++
	}
}
