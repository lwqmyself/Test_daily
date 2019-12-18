package main

import "fmt"

func main() {
	arr := []int{4, 5, 89, 5, 3, 567, 435, 43, 457, 78, 100}
	QS(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func QS(arr []int, start, end int) {
	if start >= end {
		return
	}
	tmp := arr[start]
	i := start
	j := end
	for i != j {
		for arr[j] >= tmp && i < j {
			j--
		}
		for arr[i] <= tmp && i < j {
			i++
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i], arr[start] = tmp, arr[i]
	QS(arr, start, i-1)
	QS(arr, i+1, end)

}
