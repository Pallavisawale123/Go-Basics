package main

import "fmt"

func main() {
	fmt.Println("Welcome to func in golang")

	var result = addTwoNumbers(10, 20)
	fmt.Println("Addtion of two numbers is : ", result)

	var result2 = muliTwoNumbers(2, 2, 2, 2)
	fmt.Println("Multiplication of numbers is : ", result2)

	var result3, operationType = addNumbers(8, 2, 2)

	fmt.Println("result of addition: ", result3)
	fmt.Println("values type is : ", operationType)
}

func addTwoNumbers(a int, b int) int {
	// fmt.Println("Addition of two numbers is : ", a+b)
	return a + b
}

func muliTwoNumbers(values ...int) int {

	resultOfMulti := 1

	for _, v := range values {
		resultOfMulti *= v
	}
	return resultOfMulti
}

//return 2 values from func

func addNumbers(values ...int) (int, string) {

	var result int
	for _, v := range values {
		result += v
	}
	return result, "Given operation is of type int"
}
