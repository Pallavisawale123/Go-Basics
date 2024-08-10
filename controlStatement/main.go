package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("If Else and switch in golang")

	// loginCount := 23

	// var result string

	// if loginCount < 0 {
	// 	result = "Not valid"
	// } else {
	// 	result = "valid"
	// }

	// fmt.Println("Result is : ", result)

	// if num := 2; num == 2 {
	// 	fmt.Println("nums are equal")
	// } else {
	// 	fmt.Println("nums not equal")
	// }

	//switch examples

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1

	fmt.Println("value of dice is :", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("dice value is one you can open")
	case 2:
		fmt.Println("dice value is two you can't open")
	default:
		fmt.Println("not good")

	}

}
