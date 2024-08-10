package main

import "fmt"

func main() {

	fmt.Println("break continue examples in go")

	// days := []string{"sund", "mond", "tues", "thur", "fri", "sat"}

	// for _, value := range days {
	// 	fmt.Println("values are :", value)
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	score := 1

	for score < 6 {
		if score == 4 {
			score++
			break
		}
		fmt.Println("value of score is : ", score)
		score++
	}

	score1 := 1

	fmt.Println("continue example in golang")
	for score1 < 6 {
		if score1 == 4 {
			score1++
			continue
		}
		fmt.Println("value of score is : ", score1)
		score1++
	}

}
