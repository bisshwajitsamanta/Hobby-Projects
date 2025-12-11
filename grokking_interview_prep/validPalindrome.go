package main

import (
	"fmt"
	"unicode"
)

func main() {
	inputString := "RACECAR"
	fmt.Printf("%s ::=> is a Valid Palindrome %v\n", inputString, ValidPalindrome(inputString))
}

func ValidPalindrome(inputString string) bool {

	left, right := 0, len(inputString)-1
	for left < right {

		for left < right && !unicode.IsLetter(rune(inputString[left])) && !unicode.IsDigit(rune(inputString[left])) {
			left++
		}
		for left < right && !unicode.IsLetter(rune(inputString[right])) && !unicode.IsDigit(rune(inputString[right])) {
			right--
		}
		if unicode.ToLower(rune(inputString[left])) != unicode.ToLower(rune(inputString[right])) {
			return false
		}
		left++
		right--
	}
	return true
}
