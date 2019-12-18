package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "a"

	fmt.Println(lengthOfLastWord(s))
}
func lengthOfLastWord(s string) int {
	if len(s) <= 0 {
		return 0
	}
	s = strings.Trim(s, " ")

	res := strings.Split(s, " ")
	fmt.Println(res)
	num := len(res[len(res)-1])

	return num
}
