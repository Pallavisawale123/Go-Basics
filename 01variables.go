package main

import "fmt"

func main() {

	var aa int = 4
	fmt.Println("int value is ", aa)

	var sname string = "pallavi"
	fmt.Println("Name of the student is :", sname)

	pname := "sawale"
	fmt.Println("pname is : ", pname)

	scoreRating := 3.4
	fmt.Printf("%v ,%T ", scoreRating, scoreRating)

	var pscore, Sscore int = 4, 6
	fmt.Print("pscore is ", pscore, "Sscore is ", Sscore)

	var intDefaultValue int
	fmt.Print("default value of int is ", intDefaultValue)

	var (
		spiderman = "Im spiderman"
		age       = 20
		powers    = "web"
	)

	fmt.Print(spiderman, age, powers)
}
