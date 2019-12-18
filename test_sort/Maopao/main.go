package main

import "fmt"

func main() {
	arr := []int{4, 5, 89, 5, 3, 567, 435, 43, 457, 78, 100}
	maopao(arr)
	fmt.Println(arr)
}

func maopao(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		fmt.Println(arr)
	}
}
