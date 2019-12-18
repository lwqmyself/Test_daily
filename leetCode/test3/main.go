package main

import (
	"fmt"
)

func main() {
	str := "abba"
	fmt.Println(lengthOfLongestSubstring(str))
}
func lengthOfLongestSubstring(s string) int {
	max := 0
	//fmt.Println(reflect.TypeOf(s[1]))
	m := make(map[int]int)
	str := []rune(s)
	for i, j := 0, 0; i < len(s); i++ {
		val := int(str[i])
		if index, ok := m[val]; index >= j && ok {
			j = index + 1
		}
		num := i - j + 1
		//fmt.Println(num,i,j)
		if num > max {
			max = num
		}
		m[val] = i
	}

	return max
}
