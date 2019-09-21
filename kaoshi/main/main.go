package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		n int
		m int
	)
	fmt.Scanf("%d %d", &n, &m)
	var aa []int
	sum := make([]int, 0)
	sum1 := make([]int, 0)
	for i := 0; i < n; i++ {
		tmp := 0
		fmt.Scanf("%d", &tmp)
		aa = append(aa, tmp)
	}
	//fmt.Println(aa)
	for i := m; i <= n; i++ {
		for j := 0; j < n-m+1; j++ {
			cc := 0
			for k := j; k < j+m; k++ {

				cc += aa[k]

			}
			sum1 = append(sum1, cc)
		}
		sort.Ints(sum1)
		sum = append(sum, sum1[0])
	}
	sort.Ints(sum)
	fmt.Println(sum[0])
}
