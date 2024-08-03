package main

import "fmt"

func isUnique(myStr string) bool {
	myMap := make(map[rune]bool)

	for _, val := range myStr {

		if myMap[val] == true {

			return true
		} else {

			myMap[val] = true
		}

	}
	return false
}

func main() {
	myStr := "hello"
	if isUnique(myStr) == true {
		fmt.Println("not unique")
	} else {
		fmt.Println("unique")
	}

}
