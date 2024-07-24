package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	var username string
	var userRating string

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("enter you full name")

	username, _ = reader.ReadString('\n')

	//take rating from user and convert to int

	reader = bufio.NewReader(os.Stdin)

	fmt.Printf("enter your rating")
	userRating, _ = reader.ReadString('\n')
	mynumRating, _ := strconv.ParseInt(strings.TrimSpace(userRating), 2, 64)

	fmt.Print("%v \n Thanks for rating ,%v \n", username, mynumRating, time.Now().Format(time.Stamp))

}
