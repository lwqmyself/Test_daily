package main

import (
	"fmt"
)

func main() {
	s := "fbfdc"
	Ccc(s, 2)
}

func Ccc(s string, key int) {

	res := make([]string, 0)

	for i := 0; i < len(s); i++ {
		for j := 1; j < len(s); j++ {
			if i+j > len(s) {
				break
			}
			res = append(res, s[i:i+j])
		}
	}
	fmt.Println(res)
	m := make(map[string]int)
	max := 0
	for _, v := range res {
		tmp := make(map[string]int)
		for i := 0; i < len(v); i++ {
			tmp[v[i:i+1]] = 1
		}
		fmt.Println(tmp)
		if len(tmp) == key {
			m[v] = len(v)
			if len(v) > max {
				max = len(v)
			}
		}
		fmt.Println(m)

	}

	for key, v := range m {
		if v == max {
			fmt.Println(key)
		}
	}

}
