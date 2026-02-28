// package palindrome provides a solution for the "Valid Palindrome" problem.
package palindrome

import (
	"unicode"
)

// IsPalindrome checks if a string is a palindrome, considering only alphanumeric characters and igonoring case.
// it returns true if the string is a palindrome, false otherwise.
//
// Time complexity: O(n)
// Space complexity: O(n)
func IsPalindrome(s string) bool {
	runes := []rune(s)
	left, right := 0, len(runes)-1

	for left < right {
		for left < right && !unicode.IsLetter(runes[left]) && !unicode.IsDigit(runes[left]) {
			left++
		}
		for left < right && !unicode.IsLetter(runes[right]) && !unicode.IsDigit(runes[right]) {
			right--
		}
		if unicode.ToLower(runes[left]) != unicode.ToLower(runes[right]) {
			return false
		}

		left++
		right--
	}

	return true
}
