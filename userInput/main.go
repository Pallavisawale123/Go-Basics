package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const PIEx = 3.14

func main() {

	fmt.Println("Welcome to our pizza app")
	fmt.Println("Give the rating between 1 to 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating ", input)

	numRating, err := strconv.ParseInt(strings.TrimSpace(input), 8, 32)
	fmt.Println(numRating)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("rating passed")
	}
}
