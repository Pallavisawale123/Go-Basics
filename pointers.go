package main

import "fmt"

func main() {

	var tt = 20

	var ss *int = &tt

	fmt.Println("pointer value : ", *ss)
	fmt.Println("Address of tt variable ", &tt)
	fmt.Println("Address of ss variable", &ss)

	var ss1 = 20

	pp1 := &ss1

	fmt.Println(*pp1)

}
