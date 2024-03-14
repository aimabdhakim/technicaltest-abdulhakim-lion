package main

import (
	"fmt"
	"strings"
)

func main() {
	// You can change variable below "a" or "b" to check another strings you want
	a := "radar"
	b := "hello"

	fmt.Println(palindromeCheck(a))
	fmt.Println(palindromeCheck(b))
}

func palindromeCheck(str string) bool {
	palindrome := false
	strLength := len(str)
	reverseStr := ""

	for i := strLength - 1; i >= 0; i-- {
		reverseStr = reverseStr + string(str[i])
	}

	if strings.ToLower(str) == strings.ToLower(reverseStr) {
		palindrome = true
	}

	return palindrome
}
