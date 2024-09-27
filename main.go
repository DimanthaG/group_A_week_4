package main

import "fmt"

func main() {
	fmt.Println("Welcome to Group A's Week 4 Project!")
    nums := []int{64, 34, 25, 12, 22, 11, 90} //Dimantha Goonewardena - Bubble Sort
	fmt.Println("Unsorted:", nums)
	bubbleSort(nums)
    fmt.Println("Sorted:", nums)
	fmt.Println("sum of an array is ",sumArray(nums)) //Divyanshu - Array Sum
	result, err := errorHandling(5, 0)	//Sujan Thapa - Error Handling
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
