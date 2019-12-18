package main

import "fmt"

func main() {
	arr := []int{10, 34, 19, 100, 80, 1221, 1, 3, 5, 989, 565, 8998}
	SeSort(arr)
	fmt.Println(arr)
}

func SeSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		minVal := arr[i]
		minIndex := i
		for j := minIndex; j < len(arr); j++ {
			if minVal > arr[j] {
				minVal = arr[j]
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = minVal, arr[i]
	}
}
