package main

import (
	"fmt"
	"time"
)

func main() {
	arr := []int{4, 5, 89, 5, 3, 567, 435, 43, 457, 78, 100}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func QuickSort(arr []int, start, end int) {
	if start >= end { //需要排序的起始位置大于或等于终止位置，就表明不再需要排序
		return
	}
	tmp := arr[start]
	i := start
	j := end
	fmt.Println(start, end)
	for i != j {
		for arr[j] >= tmp && i < j {
			//fmt.Println("j=", j)
			j--
		}
		for arr[i] <= tmp && i < j {
			i++
		}
		//fmt.Println(i)

		fmt.Println(i, j)
		time.Sleep(time.Second)
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]

		}

	}
	arr[i], arr[start] = tmp, arr[i]
	//递归处理基数左边未处理的
	QuickSort(arr, start, i-1)

	//递归处理基数右边未处理的
	if i != len(arr)-1 {
		QuickSort(arr, i+1, end)
	}

}
