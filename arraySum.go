package main


// Function to calculate the sum of an array
func sumArray(arr []int) int {
    sum := 0
    for _, num := range arr {
        sum += num
    }
    return sum
}