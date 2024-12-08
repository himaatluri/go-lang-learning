// In this challenge, the user enters a string and a substring. You have to print the number of times that the substring occurs in the given string. String traversal will take place from left to right, not from right to left.

// NOTE: String letters are case-sensitive.

// Input Format

// The first line of input contains the original string. The next line contains the substring.

// Constraints

// Each character in the string is an ascii character.

// Output Format

// Output the integer number indicating the total number of occurrences of the substring in the original string.

// Sample Input

// ABCDCDC
// CDC
// Sample Output

// 2
package main

import (
	"fmt"
	"strings"
)

func main() {
	text := ""
	substring := ""
	fmt.Println("Enter the String")
	fmt.Scan(&text)
	fmt.Println("Enter the Subsstring")
	fmt.Scan(&substring)

	// result := strings.Contains(text, substring)
	length := strings.Count(text, substring)

	fmt.Println(length)
}
