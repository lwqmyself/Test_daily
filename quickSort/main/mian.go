package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//cc()
	/*str := "abc bcd uit"
	arr := strings.Split(str, " ")
	fmt.Println(len(arr), arr)*/
	coins := []int{1, 2, 5, 10}
	res := test1(coins, 10000)
	fmt.Println(res)
}

func cc() {

	//var a string
	/*if err!= nil {

	}*/
	//str1 := make([]string, 5, 10)
	//str2 := make([]string, 5, 10)
	inputReader := bufio.NewReader(os.Stdin)
	str, err := inputReader.ReadString('\n')
	str1 := string(str)
	if err != nil {

	}
	n := strings.Split(str1, " ")
	for i := len(n) - 1; i >= 0; i-- {
		if i == len(n)-1 {
			fmt.Print(n[i][:len(n[i])-1] + " ")
		} else {
			fmt.Print(n[i] + " ")
		}

	}

}

func QSort(n []int) {

	for i := 1; i < len(n); i++ {
		if n[i] < n[i-1] {
			j := i - 1
			temp := n[i]
			for j >= 0 && n[j] > temp {
				n[j+1] = n[j]
				j--
			}
			n[j+1] = temp
		}

	}

}
func QSort2(tree []int) {
	for i := 1; i < len(tree); i++ {
		for j := i; j > 0 && tree[j] < tree[j-1]; j-- {
			tree[j], tree[j-1] = tree[j-1], tree[j]
		}
		fmt.Println(tree)
	}

}

func test1(coins []int, n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= n; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[n]
}
