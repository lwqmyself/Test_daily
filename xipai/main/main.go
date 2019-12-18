package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
 * @author miku
 * @date 2019/9/21
 */
func main() {
	//洗牌算法
	rand.Seed(time.Now().UnixNano())
	arr := []int{4, 6, 7, 8, 2}
	for i := len(arr) - 1; i > 0; i-- {
		index := rand.Intn(i)
		arr[i], arr[index] = arr[index], arr[i]
		/*fmt.Println("1.",rand.Intn(i),i)
		fmt.Println("2.",rand.Intn(i),i)*/

	}
	fmt.Println(arr)
	fmt.Println(FisherYatesShuffle(arr))

}
func FisherYatesShuffle(array []int) []int {
	newArray := make([]int, 0)
	l := len(array)
	for i := 0; i < l; i++ {
		p := rand.Intn(len(array)) // [0 , len(array))
		//fmt.Println(len(array))
		fmt.Printf("每次生成的随机数：%d\n", p)
		newArray = append(newArray, array[p])
		array = append(array[0:p], array[p+1:]...)
		fmt.Println(array)
	}

	return newArray
}
