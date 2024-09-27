package main

// bubbleSort sorts a slice of integers in ascending order using the Bubble Sort algorithm.
func bubbleSort(nums []int) {
    n := len(nums)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if nums[j] > nums[j+1] {
                // Swap nums[j] and nums[j+1]
                nums[j], nums[j+1] = nums[j+1], nums[j]
            }
        }
    }
}
