package main

import "fmt"

func main() {
	str := "aaabbbaaa"
	m := make(map[string]int)
	for i := 0; i < len(str); i++ {
		for j := 1; j <= len(str); j++ {
			if i+j > len(str) {
				break
			}
			//判断是不是回文
			if isPalindromic(str[i : i+j]) {
				m[str[i:i+j]] = 1
			}
		}
	}
	for k, _ := range m {
		fmt.Print(k, " ")
	}

}

//判断是不是回文
func isPalindromic(str string) bool {
	//长度为1，直接返回false
	if len(str) <= 1 {
		return false
	}
	i := 0
	j := len(str) - 1
	//前后指针向中间前进
	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}
