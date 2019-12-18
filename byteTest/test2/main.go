package main

import "fmt"

func main() {
	var k int
	res := make(map[int]map[int]int, 0)
	tm := make([]int, 0)
	sum := make(map[int]int)
	count := make([]int, 0)
	fmt.Scanf("%d\n", &k)
	for i := 0; i < k; i++ {
		var (
			n int
			m int
		)
		fmt.Scanf("%d %d", &n, &m)
		//fmt.Println(n)
		for j := 0; j < n; j++ {
			tmp := 0
			fmt.Scanf("%d", &tmp)
			tm = append(tm, tmp)
		}

		for l := 0; l < len(tm); l++ {
			if l == 0 {
				sum[0] = 0
				continue
			}

			for n1 := 1; n1 <= l; n1++ {

				shengxia := m - tm[l]
				if tm[n1] > shengxia {
					count[l]++
				}
			}

		}
		res[i] = sum

	}
	fmt.Println(res)

}
