package main

import (
	"fmt"
)

func main() {

	var (
		n int
		m int
		k int
	)

	fmt.Scanf("%d%d%d", &n, &m, &k)
	left := 1
	right := 0
	mid := 0

	i := 0

	right = n * m
	count := 0
	for left <= right {
		mid = (right) / 2
		fmt.Println(mid)

		for i = n; i > 0; i-- {
			if m <= mid/i {
				count += m * i
				break
			}
			count += mid / i

			if count < k {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

	}

	fmt.Printf("%d", left)

}
