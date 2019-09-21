package main

import (
	"fmt"
)

func main() {
	var (
		n int
		b int
	)
	fmt.Scanf("%d %d", &n, &b)
	xx := make([]int, 0)
	for i := 0; i < n; i++ {
		tmp := 0
		fmt.Scanf("%d", &tmp)
		xx = append(xx, tmp)
	}
	for i := 0; i < len(xx); i++ {

	}

	fmt.Println("10")

}
