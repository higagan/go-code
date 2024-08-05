// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func isPer(myStr1, myStr2 string) bool {
	if len(myStr1) != len(myStr2) {
		return false
	}
	storeCount := make(map[rune]int)
	for _, val := range myStr1 {
		storeCount[val]++
	}

	for _, val := range myStr2 {
		storeCount[val]--
		if storeCount[val] < 0 {
			return false
		}

	}
	return true

}

func main() {
	myStr1 := "hello"
	myStr2 := "hello"
	if isPer(myStr1, myStr2) == false {
		fmt.Println("not permutation")
	} else {
		fmt.Println("permutation")
	}

}
