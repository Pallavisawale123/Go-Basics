package main

import "fmt"

func main() {

	fmt.Println("Map examples")

	var lang = make(map[string]string)

	lang["js"] = "javascript"
	lang["RB"] = "ruby"
	lang["go"] = "golang"
	lang["py"] = "python"

	fmt.Println("map contains : ", lang)

	delete(lang, "RB")

	fmt.Println("map contains : ", lang)

	//loops in go

	for key, value := range lang {
		fmt.Printf("For key %v ,values is %v \n", key, value)
	}

	for _, value := range lang {
		fmt.Printf("For key ,values is %v \n", value)
	}
}
