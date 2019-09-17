package main

import "fmt"

func SelectSort(arr []int) {

	for i := 0; i < len(arr); i++ {
		miniIndex := i
		miniVal := arr[i]
		for j := miniIndex; j <= len(arr)-1; j++ {
			if miniVal > arr[j] {
				miniVal = arr[j]
				miniIndex = j
			}
		}
		arr[i], arr[miniIndex] = arr[miniIndex], arr[i]

	}
	fmt.Println(arr)
}

func main() {
	arr := []int{10, 34, 19, 100, 80, 1221, 1, 3, 5, 989, 565, 8998}
	SelectSort(arr)
}
