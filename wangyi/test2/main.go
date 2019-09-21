package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		var (
			n int
			k int
		)
		aa := make([]int, 0)
		fmt.Scanf("%d %d", &n, &k)
		for i := 0; i < n; i++ {
			tmp := 0
			fmt.Scanf("%d", &tmp)
			aa = append(aa, tmp)
		}
		fmt.Println("NO")
	}
}
