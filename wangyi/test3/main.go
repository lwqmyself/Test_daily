package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)
	xx := make([]int, 0)
	for i := 0; i < t; i++ {

		count := 0

		var (
			a int
			b int
			p int
			q int
		)

		fmt.Scanf("%d %d %d %d", &a, &b, &p, &q)
		for {
			count++
			if a+p >= b {
				break
			}
			p = p * q
		}
		xx = append(xx, count)

	}
	for _, v := range xx {
		fmt.Println(v)
	}
}
