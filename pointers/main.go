package main

import "fmt"

func main() {

	fmt.Println("Pointers example")

	var one int = 2
	var ptr *int = &one

	var ptr1 = &ptr

	fmt.Println(&ptr)
	fmt.Println(&one)

	fmt.Println("Value of ptr is : ", *ptr)    //2
	fmt.Println("value of one is : ", one)     // 2
	fmt.Println("value of ptr1 is : ", **ptr1) // 2

}
