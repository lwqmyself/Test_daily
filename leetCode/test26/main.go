package main

import "fmt"

/*nums = [0,0,1,1,1,2,2,3,3,4],

函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。

你不需要考虑数组中超出新长度后面的元素。
说明:

为什么返回数值是整数，但输出的答案是数组呢?

请注意，输入数组是以“引用”方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。

你可以想象内部操作如下:

// nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
int len = removeDuplicates(nums);

// 在函数里修改输入数组对于调用者是可见的。
// 根据你的函数返回的长度, 它会打印出数组中该长度范围内的所有元素。*/

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	removeDuplicates(nums)
	fmt.Println(nums)
}
func removeDuplicates(nums []int) int {
	l := len(nums)
	k := 0
	if l == 0 {
		return 0
	}
	for i := 1; i < l; i++ {
		if nums[i] != nums[k] {
			//如果下标 i 和 k 的值不一样就把下标为 i 的值赋值给 k的下一个
			k++
			nums[k] = nums[i]
		}
	}
	return k + 1
}
