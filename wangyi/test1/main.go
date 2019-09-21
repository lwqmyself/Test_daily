package test1

import (
	"fmt"
)

func main() {
	var (
		n int
		m int
	)
	fmt.Scanf("%d %d", &n, &m)
	xx := make([]int, 0)
	sum := make(map[int]int)
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scanf("%d", &tmp)
		xx = append(xx, tmp)
		sum[tmp]++
	}
	//fmt.Println(xx)
	//fmt.Println(sum[6])
	count := make([]int, 0)
	tmp := 0
	for i := 0; i < m; i++ {

		fmt.Scanf("%d", &tmp)
		//fmt.Println(tmp)
		count = append(count, sum[tmp])
	}
	//fmt.Println(count)
	for _, v := range count {
		fmt.Println(v)
	}

}
