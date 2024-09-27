package main

import "fmt"

func main() {
	fmt.Println("Welcome to Group A's Week 4 Project!")
	nums := []int{64, 34, 25, 12, 22, 11, 90} //Dimantha Goonewardena - Bubble Sort
	fmt.Println("Unsorted:", nums)
	bubbleSort(nums)
	fmt.Println("Sorted:", nums)
	fmt.Println("sum of an array is ", sumArray(nums)) //Divyanshu - Array Sum
	result, err := errorHandling(5, 0)                 //Sujan Thapa - Error Handling
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	var country1, country2 string

	// Get first country from user
	fmt.Println("Enter the first country:")
	fmt.Scanln(&country1)

	// Get second country from user
	fmt.Println("Enter the second country:")
	fmt.Scanln(&country2)

	// Get the population of both countries
	population1 := compelif(country1)
	population2 := compelif(country2)

	// Compare populations
	if population1 == 0 || population2 == 0 {
		fmt.Println("Sorry, one of the countries is not in our list.")
	} else if population1 > population2 {
		fmt.Printf("%s has a larger population than %s.\n", country1, country2)
	} else if population1 < population2 {
		fmt.Printf("%s has a smaller population than %s.\n", country1, country2)
	} else {
		fmt.Printf("%s and %s have the same population.\n", country1, country2)
	}

	// Call the reverseString function
	name := "Jaspal"
	fmt.Println("The reverse of ", name, "is:", reverseString(name))
}
