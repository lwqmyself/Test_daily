package main

import "fmt"

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	flag := false

	m1, m2 := 0, 0
	length := len(nums1) + len(nums2)
	if length%2 == 0 {
		flag = true
		m1 = length / 2
		m2 = m1 + 1
	} else {
		m1 = length/2 + 1
		m2 = m1
	}
	//判断数组的长度
	alen := len(nums1)
	blen := len(nums2)
	cl := alen + blen

	//fmt.Println(cl)
	//var c [cl]int // non-constant array bound cl
	c := make([]int, cl)

	//fmt.Println(len(c))
	//fmt.Println(cap(c))
	i := 0
	j := 0
	k := 0

	for i < alen && j < blen {

		if nums1[i] < nums2[j] {
			c[k] = nums1[i]

			i++
			k++
		} else {
			c[k] = nums2[j]
			k++
			j++
		}
		fmt.Println(c)
	}

	for i < alen {
		c[k] = nums1[i]
		k++
		i++
	}
	for j < blen {
		c[k] = nums2[j]
		j++
		k++
	}

	if flag {
		return float64(c[m1-1]+c[m2-1]) / 2.0
	} else {
		return float64(c[m1-1])
	}
}

func Judge(flag bool, k int) bool {

	return false
}
