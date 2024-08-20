package main

import (
	"fmt"
)

func isPalindrome(s string) bool {

	n := len(s)

	for i := 0; i <= n/2; i++ {

		if s[i] != s[n-i-1] {
			return false
		}
	}
	return true

}

func main() {

	if isPalindrome("popa") {

		fmt.Println("Palindrom")

	} else {
		fmt.Println(" not Palindrom")
	}

}
