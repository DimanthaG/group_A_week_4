package main

// Function to reverse a string
func reverseString(s string) string {
	// Convert the string to a slice of reverse
	reverse := []rune(s)
	for i, j := 0, len(reverse)-1; i < j; i, j = i+1, j-1 {
		// Swap the characters at the beginning and end of the slice
		reverse[i], reverse[j] = reverse[j], reverse[i]
	}
	return string(reverse) // Convert the slice of reverse back to a string
}
