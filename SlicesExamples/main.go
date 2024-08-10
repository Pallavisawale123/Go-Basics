package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices Examples")

	var fruitList = []string{"m1", "m2", "m3", "m4"}

	fmt.Println("fruitList contains : ", fruitList)

	fruitList = append(fruitList, "m5")

	fmt.Println("fruitList contains : ", fruitList)

	fmt.Println("fruitList contains : ", fruitList[1:2])
	fmt.Println("fruitList contains : ", fruitList[:2])

	var scores = make([]int, 4)

	fmt.Println("Print the scores : ", scores)

	scores[0] = 5
	scores[1] = 15
	scores[2] = 25
	scores[3] = 35
	// scores[4] = 45

	scores = append(scores, 45, 55, 65, 75)
	fmt.Println("Print the scores : ", scores)

	//Sort the array

	sort.Ints(scores)
	fmt.Println("Print the scores : ", scores)
	fmt.Println("Print the scores : ", sort.IntsAreSorted(scores))

}
