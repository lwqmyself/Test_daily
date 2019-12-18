package main

import "fmt"

func main() {

}

func AddNums(nums []int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	fmt.Println(sum)
}
