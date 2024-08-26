package main

import (
	"fmt"
)

// reverseString function to reverse the input string
func reverseString(s string) string {
	// Convert string to rune slice to handle Unicode characters
	runes := []rune(s)
	// Reverse the slice
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	// Convert back to string
	return string(runes)
}

func main() {
	var input string
	fmt.Print("Enter a string to reverse: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	reversed := reverseString(input)
	fmt.Println("Reversed:", reversed)
}
