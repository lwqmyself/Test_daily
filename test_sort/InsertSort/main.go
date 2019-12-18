package main

import "fmt"

func main() {
	arr := []int{4, 5, 89, 5, 3, 567, 435, 43, 457, 78, 100}
	InsertSort(arr)

	fmt.Println(arr)
}
func InsertSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}
