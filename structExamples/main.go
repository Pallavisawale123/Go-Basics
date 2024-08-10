package main

import "fmt"

func main() {

	fmt.Println("struct examples")

	hitesh := User{"hitesh", "hitesh@gmail.com", true, 20}
	fmt.Println(hitesh)

	fmt.Printf("hitesh details are :%+v", hitesh)

	fmt.Printf("hitesh name  is : %v and email is : %v ", hitesh.Name, hitesh.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
