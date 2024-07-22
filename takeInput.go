package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// var name string
	// fmt.Scanln(&name)
	// fmt.Println("Your name is : ", name)

	// var name1 = "pallavi"
	// var a, _ = fmt.Println(name1)
	// fmt.Println(a)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter full name ")
	// myname, _ := reader.ReadString('\n')
	// fmt.Print(myname)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter number")

	myrating, _ := reader.ReadString('\n')

	muNumRating, _ := strconv.ParseFloat(strings.TrimSpace(myrating), 64)
	fmt.Println(muNumRating)

}
