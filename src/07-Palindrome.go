// Go Problem 7 - Palindrome Test
// Hugh Brady - G00338260 - 05/10/17

package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	
	// prompt and read user input
	fmt.Println("Enter String: ")
	fmt.Scanf("%s\n", &input)
	
	input = strings.ToLower(input)			// convert to lowercase to ensure palindrome checker is accurate
	fmt.Println(isPal(input))				// call function to check if input is a palindrome
}

// function to check if string is a palindrome
func isPal(s string) string {
	mid := len(s) / 2
	last := len(s) - 1
	
	// This for loop compares the first and last letters in the string, then moves in and compares the 
	// second and second last. This continues until the string is proven not to be a palindrome
	for i := 0; i < mid; i++ {
		if s[i] != s[last-i] {
			return "No, this is not a palindrome"
		}
	}
	
	// If we haven't broken into the if statement above that means we have a palindrome!
	return "Yes, this is a palindrome!"
}
