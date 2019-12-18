package main

import (
	"fmt"
	"strconv"
)

func main() {
	//
	fmt.Println(countAndSay(5))
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	if n == 2 {
		return "11"
	}
	p := countAndSay(n - 1)
	//fmt.Println("p=",p)
	count := 1
	pre := StringToRuneArr(p)
	//fmt.Println("pre = ",pre)
	res := make([]string, 0)
	for i := 1; i < len(pre); i++ {
		//fmt.Println(reflect.TypeOf(pre[i-1]))
		if pre[i] == pre[i-1] {
			count++
			continue
		} else {
			res = append(res, strconv.Itoa(count))
			res = append(res, fmt.Sprintf("%v", pre[i-1]))
			count = 1

		}
	}
	res = append(res, strconv.Itoa(count))
	res = append(res, string(pre[len(pre)-1]))
	//fmt.Println(res)
	str := SliceToString(res)

	return str
}
func SliceToString(s []string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		res += s[i]
	}
	return res
}
func StringToRuneArr(s string) []string {
	res := make([]string, 0)
	for i := 0; i < len(s); i++ {
		res = append(res, s[i:i+1])
	}
	return res
}
