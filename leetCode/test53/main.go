package main

func main() {

	nums := []int{5, 5, -5, 6, 65, -12}
	l := len(nums)

	ans := nums[0]
	result := nums[0]

	for i := 1; i < l; i++ {
		if nums[i] > nums[i]+ans {
			ans = nums[i]
		} else {
			ans = nums[i] + ans
		}
		if ans > result {
			result = ans
		}
	}

}
