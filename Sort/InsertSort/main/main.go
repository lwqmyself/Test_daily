package main

import (
	"fmt"
	"time"

	"math/rand"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	tree := make([]int, 10)
	for i := 0; i < 10; i++ {
		//tree = append(tree, )
		tree[i] = int(r.Intn(1000))
	}
	for i := 10; i < 100; i++ {
		tree = append(tree, int(r.Intn(1000)))
	}

	//insertSoet1(tree)
	//insertSort2(tree)
	insertSort3(tree)
}

func insertSoet1(array []int) {

	var i, j int
	var tmp int
	for i = 1; i < len(array); i++ {
		tmp = array[i]
		for j = i; j > 0 && array[j-1] > tmp; j-- {
			array[j] = array[j-1]
		}
		array[j] = tmp
	}
	fmt.Println(array)

}

func insertSort2(tree []int) {

	fmt.Println(tree)

	for i := 1; i < len(tree); i++ {
		for j := i; j > 0 && tree[j] < tree[j-1]; j-- {
			tree[j], tree[j-1] = tree[j-1], tree[j]
		}

	}
	fmt.Println(tree)
}

func insertSort3(array []int) {

	for i := 1; i < len(array); i++ {
		var j int
		tmp := array[i]
		for j = i; j > 0 && tmp < array[j-1]; j-- {
			array[j] = array[j-1]

		}
		array[j] = tmp

	}
	fmt.Println(array)

}
