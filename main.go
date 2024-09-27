package main

import "fmt"

func main() {
	fmt.Println("Welcome to Group A's Week 4 Project!")

	result, err := errorHandling(5, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
