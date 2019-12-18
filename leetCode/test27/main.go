package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	i := removeElement(nums, 2)
	fmt.Println(nums, i)
}
func removeElement(nums []int, val int) int {
	l := len(nums)
	//i := 0
	tmp := 0
	for i, j := 0, 0; j < l; j++ {
		fmt.Println(j)
		if nums[j] != val {
			//fmt.Println(nums[j])
			nums[i] = nums[j]
			i++
			tmp = i
		}

	}

	return tmp
}
