package main

// package conversion

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println("formated date is : ", presentTime.Format("6-8-2024 9:58:07 Tuesday"))
}
