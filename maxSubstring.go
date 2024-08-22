// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "my name is gagan"
	str_split := strings.Fields(str)
	max_len := len(str_split[0])
	myStr := str_split[0]
	for _, val := range str_split {

		if len(val) > max_len {
			max_len = len(val)
			myStr = val
		}

	}
	fmt.Println(myStr)

}
