package main

import "fmt"

func main() {
	var (
		n int
		m int
	)
	count := make(map[int]int)
	aMap := make(map[int]int)
	//bMap := make(map[int]int)
	var a int
	var b int

	fmt.Scanf("%d %d", &n, &m)
	for i := 1; i <= m; i++ {
		k := 0
		j := 0
		fmt.Scanf("%d %d", &k, &j)
		count[k] = j

	}
	//fmt.Println(count)
	sum := 0
	sum1 := 0
	for i := 1; i <= n; i++ {
		if count[i] == 0 {
			if sum%2 == 0 {

				aMap[i]++
			} else {
				b++
			}
			sum++
		} else {
			if count[i] == 1 {
				if sum1%2 == 0 {

				} else {
					b++
				}
				sum1++
			}
		}
	}
	if a > b {
		fmt.Println(b * 2)
	} else {
		fmt.Println(a * 2)
	}
}
