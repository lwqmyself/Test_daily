package main

import "fmt"

func main() {
	str := "abcabcaabca"
	sub := "abca"
	fmt.Println(count(str, sub))
}

func count(str, sub string) int {
	ls := len(str)
	lb := len(sub)
	count := 0
	for i := 0; i < ls-lb; i++ {
		if str[i:i+lb] == sub {
			count++
		}
	}
	return count
}
