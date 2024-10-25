// 125. Valid Palindrome
 //  https://leetcode.com/problems/valid-palindrome/

package main

import "fmt"

func isPalindrome(s string) bool {
	if len(s) == 1 || len(s)==2 {
		return true // example " "
	}
	var str string

	for _, char := range s {
		if char >= 'A' && char <= 'Z' || char >= 'a' && char <= 'z' || char >= '0' && char < '9' {
			if char >= 'A' && char <= 'Z' {
				char = char + 32
			}

			str += string(char)
		}
	}

	///pallinddrome solving logic
	left, right := 0, len(str)-1

	for left < right {

		if s[left] == s[right] {
			return true
		}
		left++
		right--
	}

	return false
}

func main() {
	input := "amanaplanacanalpanama    ADDD "

	fmt.Println(isPalindrome(input))
}
