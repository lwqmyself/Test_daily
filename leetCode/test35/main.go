package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, 6))
}
func searchInsert(nums []int, target int) int {
	l := len(nums)
	count := -1
	for i := 0; i < l; i++ {
		if nums[i] >= target {
			break
		}
		count = i
	}
	//fmt.Println(count)

	return count + 1

}
