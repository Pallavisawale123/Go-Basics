package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Demo API for Get request")
	PerformGetRequest("https://jsonplaceholder.typicode.com/posts")
}

func PerformGetRequest(URL string) {

	response, err := http.Get(URL)

	if err != nil {
		panic(err)
	}

	// defer response.Body.Close()

	fmt.Println("status code : {}", response.StatusCode)
	fmt.Println("Content length : {}", response.ContentLength)

	var responseString strings.Builder

	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("bytecount is : ", byteCount)

	fmt.Println(responseString.String())
}
