package main

import (
	"fmt"
)

func main() {
	var (
		n int
		m int
	)
	fmt.Scanf("%d %d", &n, &m)
	var aa []int
	//	sum := make([]int, 0)
	res1 := make([]float64, 0)
	for i := 0; i < n; i++ {
		tmp := 0
		fmt.Scanf("%d", &tmp)
		aa = append(aa, tmp)
	}
	//fmt.Println(aa)
	//max :=0
	for i := m; i <= n; i++ {
		sum1 := make([]int, 0)
		for j := 0; j < n-i+1; j++ {
			cc := 0
			for k := j; k < j+i; k++ {

				cc += aa[k]

			}
			sum1 = append(sum1, cc)
			//fmt.Println(sum1,i)
		}
		//fmt.Println(sum1,i)
		max := 0
		l1 := len(sum1)
		for i := 0; i < l1; i++ {
			if sum1[i] > max {
				max = sum1[i]
			}
		}
		//sum = append(sum, sum1[len(sum1)-1])
		res1 = append(res1, float64(max)/float64(i))
		fmt.Println(res1)
	}
	l := len(res1)
	max := 0.0
	for i := 0; i < l; i++ {
		if res1[i] > max {
			max = res1[i]
		}
	}
	fmt.Printf("%.3f", max)
}
