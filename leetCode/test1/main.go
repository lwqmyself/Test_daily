package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 2, 4}
	res := twoSum(nums, 6)
	fmt.Println(res)
}
func twoSum(nums []int, target int) []int {
	/*res := make([]int,0)
	flag := false
	for i:=0;i<len(nums) ;i++  {
		for j:=i+1;j<len(nums) ;j++  {
			//fmt.Println(i,j)
			if nums[i]+nums[j] ==target {
				res=append(res,i,j)
				flag = true
				//fmt.Println(res)
				break
			}

		}
		if flag {
			break
		}
		//break
	}
	return res*/
	m := make(map[int]int)
	for k, v := range nums {
		//去map里面找是不是有对应的
		//就是
		k1, ok := m[v]
		if ok {
			return []int{k1, k}
		}
		//当前第 k 个元素需要 target-v 来凑成target
		//比如 nums := []int{3, 2, 4}  target = 6
		//m[3] = 0 m[4] = 1 m[2] = 2
		//第0个数需要 3 来组成 6
		//第1个数需要 4 来组成 6
		//第2个数需要 2 来组成 6
		//遍历到 第2个数的时候去map里面找它自身是不是被其他元素需要
		m[target-v] = k
	}
	return nil
}
